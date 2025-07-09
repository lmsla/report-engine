package services

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"os/exec"
	"report-backend-golang/global"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

var jwtSecret = []byte("your-secret-key") // 用於簽名 JWT 的密鑰
var blacklistedTokens = make(map[string]time.Time)

type UserInfo struct {
	UserName  string
	UserPass  string
	UserGroup string
	Exp       time.Time
}

// TokenInfo 儲存 token 和過期時間
type TokenInfo struct {
	Token     string
	ExpiresAt time.Time
}

type TokenManager struct {
	tokens map[string]TokenInfo
	mu     sync.RWMutex
}

func NewTokenManager() *TokenManager {
	return &TokenManager{
		tokens: make(map[string]TokenInfo),
	}
}

func parseDuration(s string) (time.Duration, error) {
	if s == "" {
		return time.Hour * 8, nil // 預設 8 小時
	}

	unit := s[len(s)-1]
	val := s[:len(s)-1]

	n, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("invalid duration: %s", s)
	}

	switch unit {
	case 'm':
		return time.Minute * time.Duration(n), nil
	case 'h':
		return time.Hour * time.Duration(n), nil
	case 'd':
		return time.Hour * 24 * time.Duration(n), nil
	default:
		return 0, fmt.Errorf("unsupported time unit: %c", unit)
	}
}

// GenerateToken 產生 JWT token
func GenerateToken(userName, userPass, userRole string) (string, error) {
	duration, err := parseDuration(global.EnvConfig.Auth.Radius.TokenLifespan)
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(duration)

	claims := jwt.MapClaims{
		"name": userName,
		"pass": userPass,
		"role": userRole,
		"exp":  expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateTokenAndGetUserInfo(tokenStr string) (*UserInfo, error) {
	if IsBlacklisted(tokenStr) {
		return nil, fmt.Errorf("invalid token")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 確認簽名演算法是否為 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims structure")
	}

	expUnix, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid expiration time")
	}
	expTime := time.Unix(int64(expUnix), 0)
	if time.Now().After(expTime) {
		return nil, fmt.Errorf("token expired")
	}

	name, _ := claims["name"].(string)
	pass, _ := claims["pass"].(string)
	group, _ := claims["group"].(string)

	return &UserInfo{
		UserName:  name,
		UserPass:  pass,
		UserGroup: group,
		Exp:       expTime,
	}, nil
}

func RadiusAuthenticate(username, password string) (map[string]interface{}, error) {
	// 取得 RADIUS 配置
	server := global.EnvConfig.Auth.Radius.Server
	nasPort := global.EnvConfig.Auth.Radius.NasPort
	secret := global.EnvConfig.Auth.Radius.Secret
	authMethod := global.EnvConfig.Auth.Radius.AuthMethod

	if server == "" || nasPort == "" || secret == "" {
		return nil, fmt.Errorf("radius configuration incomplete")
	}

	// 設定預設認證方式
	if authMethod == "" {
		authMethod = "radtest"
	}

	fmt.Printf("RADIUS 認證 - 用戶: %s, 伺服器: %s:%s, 方式: %s\n", username, server, nasPort, authMethod)

	var radiusOutput string
	var err error

	// 根據配置選擇認證方式
	switch authMethod {
	case "radtest":
		radiusOutput, err = authenticateWithRadtest(username, password, server, nasPort, secret)
	case "layeh":
		radiusOutput, err = authenticateWithLayeh(username, password, server, nasPort, secret)
	default:
		return nil, fmt.Errorf("unknown auth method: %s", authMethod)
	}
	fmt.Println("radiusOutput:", radiusOutput)

	if err != nil {
		return nil, err
	}

	// 統一的 group 驗證
	groupValue, err := validateRadiusGroups(radiusOutput)
	if err != nil {
		return nil, fmt.Errorf("radius group validation failed: %v", err)
	}

	// 產生 JWT token
	token, err := GenerateToken(username, password, groupValue)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %v", err)
	}

	// 回傳結果
	result := map[string]interface{}{
		"username": username,
		"source":   fmt.Sprintf("radius-%s", authMethod),
		"server":   server,
		"group":    groupValue,
		"token":    token,
	}

	return result, nil
}

func authenticateWithRadtest(username, password, server, nasPort, secret string) (string, error) {
	radtestPath := global.EnvConfig.Auth.Radius.RadtestPath
	if radtestPath == "" {
		radtestPath = "radtest"
	}

	fmt.Printf("使用 radtest 執行認證\n")

	// 執行 radtest 命令
	cmd := exec.Command(radtestPath, username, password, server, nasPort, secret)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("radtest 執行失敗: %v\nStderr: %s\n", err, stderr.String())

		if strings.Contains(err.Error(), "executable file not found") {
			return "", fmt.Errorf("radtest not found at path '%s'. Please install freeradius-utils or update radtest_path in config", radtestPath)
		}

		return "", fmt.Errorf("radius authentication failed: invalid credentials")
	}

	output := out.String()
	fmt.Printf("radtest 回應: %s\n", output)

	// 檢查是否認證成功
	if !strings.Contains(output, "Access-Accept") {
		return "", fmt.Errorf("radius authentication failed: access denied")
	}

	return output, nil
}

func authenticateWithLayeh(username, password, server, nasPort, secret string) (string, error) {
	fmt.Printf("=== layeh/radius 認證開始 ===\n")
	fmt.Printf("Server: %s\n", server)
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("NAS Port: %s\n", nasPort)

	// ... 設定邏輯 ...

	// 解析配置
	timeout, err := time.ParseDuration(global.EnvConfig.Auth.Radius.Layeh.Timeout)
	if err != nil {
		timeout = 5 * time.Second // 預設 5 秒
	}

	// 使用 127.0.0.1 作為 NAS-IP-Address (與 radtest 一致)
	nasIPAddress := "127.0.0.1"

	// 可選：允許從配置覆蓋
	if global.EnvConfig.Auth.Radius.Layeh.NasIPAddress != "" {
		nasIPAddress = global.EnvConfig.Auth.Radius.Layeh.NasIPAddress
	}

	fmt.Printf("NAS IP Address: %s\n", nasIPAddress)
	fmt.Printf("Timeout: %v\n", timeout)

	// 建立 RADIUS 封包
	packet := radius.New(radius.CodeAccessRequest, []byte(secret))

	// 添加屬性並記錄 - 按照 radtest 的順序和方式
	if err := rfc2865.UserName_AddString(packet, username); err != nil {
		return "", fmt.Errorf("failed to add username: %v", err)
	}
	fmt.Printf("✓ Username 屬性已添加\n")

	if err := rfc2865.UserPassword_AddString(packet, password); err != nil {
		return "", fmt.Errorf("failed to add password: %v", err)
	}
	fmt.Printf("✓ User-Password 屬性已添加\n")

	if err := rfc2865.NASIPAddress_Add(packet, net.ParseIP(nasIPAddress)); err != nil {
		return "", fmt.Errorf("failed to add NAS IP: %v", err)
	}
	fmt.Printf("✓ NAS-IP-Address 屬性已添加: %s\n", nasIPAddress)

	// NAS-Port
	if nasPortNum, err := strconv.ParseUint(nasPort, 10, 32); err == nil {
		if err := rfc2865.NASPort_Add(packet, rfc2865.NASPort(nasPortNum)); err != nil {
			fmt.Printf("⚠️  NAS-Port 添加失敗: %v\n", err)
		} else {
			fmt.Printf("✓ NAS-Port 屬性已添加: %s\n", nasPort)
		}
	}

	// Message-Authenticator 會由 layeh/radius 自動處理
	fmt.Printf("✓ 所有基本屬性已添加，準備發送請求\n")

	// 調試：打印即將發送的請求屬性
	fmt.Printf("🔍 layeh/radius 發送的請求屬性:\n")
	for i, attr := range packet.Attributes {
		fmt.Printf("  請求屬性 %d: Type=%d, Data=%s\n",
			i+1, attr.Type, hex.EncodeToString(attr.Attribute))
	}

	// 發送請求 - 檢查 server 是否已包含端口號
	radiusAddr := server
	if !strings.Contains(server, ":") {
		radiusAddr = server + ":1812"
	}
	fmt.Printf("發送 RADIUS 請求到: %s\n", radiusAddr)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	response, err := radius.Exchange(ctx, packet, radiusAddr)
	if err != nil {
		fmt.Printf("❌ RADIUS 請求失敗: %v\n", err)
		return "", fmt.Errorf("radius exchange failed: %v", err)
	}

	fmt.Printf("✓ 收到 RADIUS 回應\n")
	fmt.Printf("Response Code: %d\n", response.Code)
	fmt.Printf("Response Identifier: %d\n", response.Identifier)

	// 檢查回應代碼
	if response.Code == radius.CodeAccessReject {
		fmt.Printf("❌ 認證被拒絕 (Access-Reject)\n")
		return "", fmt.Errorf("radius authentication failed: access rejected")
	} else if response.Code != radius.CodeAccessAccept {
		fmt.Printf("❌ 未預期的回應代碼: %d\n", response.Code)
		return "", fmt.Errorf("radius authentication failed: unexpected response code %d", response.Code)
	}

	fmt.Printf("✅ layeh/radius 認證成功 (Access-Accept)\n")
	fmt.Printf("=== layeh/radius 認證結束 ===\n")
	fmt.Println("response:", response)
	// 將 layeh 的回應轉換成與 radtest 相同的格式
	return convertLayehToRadtestFormat(response, server, nasIPAddress), nil
}

// func authenticateWithLayeh(username, password, server, nasPort, secret string) (string, error) {
// 	fmt.Printf("使用 layeh/radius 執行認證\n")

// 	// 解析配置
// 	timeout, err := time.ParseDuration(global.EnvConfig.Auth.Radius.Layeh.Timeout)
// 	if err != nil {
// 		timeout = 5 * time.Second // 預設 5 秒
// 	}

// 	nasIPAddress := global.EnvConfig.Auth.Radius.Layeh.NasIPAddress
// 	if nasIPAddress == "" {
// 		nasIPAddress = "127.0.0.1" // 預設值
// 	}

// 	nasIdentifier := global.EnvConfig.Auth.Radius.Layeh.NasIdentifier
// 	if nasIdentifier == "" {
// 		nasIdentifier = "report-backend" // 預設值
// 	}

// 	// 建立 RADIUS 封包
// 	packet := radius.New(radius.CodeAccessRequest, []byte(secret))

// 	// 添加基本屬性
// 	if err := rfc2865.UserName_AddString(packet, username); err != nil {
// 		return "", fmt.Errorf("failed to add username: %v", err)
// 	}

// 	if err := rfc2865.UserPassword_AddString(packet, password); err != nil {
// 		return "", fmt.Errorf("failed to add password: %v", err)
// 	}

// 	if err := rfc2865.NASIPAddress_Add(packet, net.ParseIP(nasIPAddress)); err != nil {
// 		return "", fmt.Errorf("failed to add NAS IP: %v", err)
// 	}

// 	if err := rfc2865.NASIdentifier_AddString(packet, nasIdentifier); err != nil {
// 		return "", fmt.Errorf("failed to add NAS identifier: %v", err)
// 	}

// 	// NAS-Port (將字串轉為數字)
// 	if nasPortNum, err := strconv.ParseUint(nasPort, 10, 32); err == nil {
// 		if err := rfc2865.NASPort_Add(packet, rfc2865.NASPort(nasPortNum)); err != nil {
// 			return "", fmt.Errorf("failed to add NAS port: %v", err)
// 		}
// 	}

// 	// 建立客戶端並發送請求
// 	ctx, cancel := context.WithTimeout(context.Background(), timeout)
// 	defer cancel()

// 	response, err := radius.Exchange(ctx, packet, server+":1812")
// 	if err != nil {
// 		return "", fmt.Errorf("radius exchange failed: %v", err)
// 	}

// 	// 檢查回應代碼
// 	if response.Code != radius.CodeAccessAccept {
// 		return "", fmt.Errorf("radius authentication failed: access denied (code: %d)", response.Code)
// 	}

// 	fmt.Printf("layeh/radius 認證成功\n")

// 	// 將 layeh 的回應轉換成與 radtest 相同的格式
// 	return convertLayehToRadtestFormat(response, server, nasIPAddress), nil
// }

func convertLayehToRadtestFormat(response *radius.Packet, server, nasIP string) string {
	var output strings.Builder

	// 處理 server 地址，避免重複端口號
	serverAddr := server
	if strings.Contains(server, ":") {
		// 已包含端口號，提取 IP 部分
		parts := strings.Split(server, ":")
		serverAddr = parts[0]
	}

	// 模擬 radtest 的輸出格式
	output.WriteString(fmt.Sprintf("Received Access-Accept Id %d from %s:1812 to %s:33859\n",
		response.Identifier, serverAddr, nasIP))

	// 調試：打印所有接收到的屬性
	fmt.Printf("🔍 layeh/radius 收到的所有屬性:\n")
	for i, avp := range response.Attributes {
		fmt.Printf("  屬性 %d: Type=%d, Length=%d, Data=%s\n",
			i+1, avp.Type, len(avp.Attribute), hex.EncodeToString(avp.Attribute))
	}

	// 解析並格式化屬性
	for _, avp := range response.Attributes {
		switch avp.Type {
		case 80: // Message-Authenticator
			output.WriteString(fmt.Sprintf("        Message-Authenticator = 0x%s\n", hex.EncodeToString(avp.Attribute)))
		case 26: // VSA (Vendor-Specific Attributes)
			fmt.Printf("🔍 發現 VSA 屬性 (Type 26), 長度: %d\n", len(avp.Attribute))
			if len(avp.Attribute) >= 6 {
				vendorID := binary.BigEndian.Uint32(avp.Attribute[0:4])
				vsaType := avp.Attribute[4]
				vsaLength := avp.Attribute[5]

				fmt.Printf("  Vendor ID: %d, VSA Type: %d, VSA Length: %d\n", vendorID, vsaType, vsaLength)

				if int(vsaLength) <= len(avp.Attribute)-4 && vsaLength >= 2 {
					vsaData := avp.Attribute[6 : 4+int(vsaLength)]
					fmt.Printf("  VSA Data: %s (hex: %s)\n", string(vsaData), hex.EncodeToString(vsaData))

					// 檢查 Fortinet VSA
					if vendorID == 12356 { // Fortinet vendor ID
						fmt.Printf("✅ 發現 Fortinet VSA\n")
						if vsaType == 35 { // Policy group type (實際是 35，不是 1)
							policyGroup := string(vsaData)
							fmt.Println("policyGroup:", policyGroup)
							output.WriteString(fmt.Sprintf("        Fortinet-FDD-SPP-Policy-Group = \"%s\"\n", policyGroup))
							fmt.Printf("✅ 解析出 Policy Group: %s\n", policyGroup)
						} else {
							fmt.Printf("⚠️  未知的 Fortinet VSA Type: %d\n", vsaType)
						}
					} else {
						fmt.Printf("⚠️  非 Fortinet VSA (Vendor ID: %d)\n", vendorID)
					}
				} else {
					fmt.Printf("❌ VSA 長度異常: declared=%d, available=%d\n", vsaLength, len(avp.Attribute)-4)
				}
			} else {
				fmt.Printf("❌ VSA 屬性太短: %d bytes\n", len(avp.Attribute))
			}
		default:
			// 其他未知屬性
			fmt.Printf("🔍 未知屬性 Type %d: %s\n", avp.Type, hex.EncodeToString(avp.Attribute))
		}
	}
	return output.String()
}

// RadiusAuthenticate 真實的 RADIUS 認證函數
// func RadiusAuthenticate(username, password string) (map[string]interface{}, error) {
// 	// 取得 RADIUS 配置
// 	server := global.EnvConfig.Auth.Radius.Server
// 	nasPort := global.EnvConfig.Auth.Radius.NasPort
// 	secret := global.EnvConfig.Auth.Radius.Secret
// 	authMethod := global.EnvConfig.Auth.Radius.AuthMethod
// 	radtestPath := global.EnvConfig.Auth.Radius.RadtestPath

// 	if server == "" || nasPort == "" || secret == "" {
// 		return nil, fmt.Errorf("radius configuration incomplete")
// 	}

// 	fmt.Printf("RADIUS 真實認證 - 用戶: %s, 伺服器: %s:%s\n", username, server, nasPort)

// 	// 檢查 radtest 路徑
// 	if radtestPath == "" {
// 		radtestPath = "radtest" // 預設值，依賴 PATH
// 	}

// 	// 執行 radtest 命令
// 	cmd := exec.Command(radtestPath, username, password, server, nasPort, secret)

// 	// 執行 radtest 命令
// 	// cmd := exec.Command("radtest", username, password, server, nasPort, secret)

// 	var out, stderr bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Printf("RADIUS 認證失敗: %v\nStderr: %s\n", err, stderr.String())
// 		return nil, fmt.Errorf("radius authentication failed: invalid credentials")
// 	}

// 	output := out.String()
// 	fmt.Printf("RADIUS 回應: %s\n", output)

// 	// 檢查是否認證成功
// 	if !strings.Contains(output, "Access-Accept") {
// 		return nil, fmt.Errorf("radius authentication failed: access denied")
// 	}

// 	// 新增 group 驗證
// 	var groupValue string

// 	groupValue, err = validateRadiusGroups(output)
// 	if err != nil {
// 		return nil, fmt.Errorf("radius group validation failed: %v", err)
// 	}
// 	// 產生 JWT token
// 	token, err := GenerateToken(username, password, groupValue)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate token: %v", err)
// 	}
// 	// 回傳結果
// 	result := map[string]interface{}{
// 		"username": username,
// 		"source":   "radius",
// 		"server":   server,
// 		"group":    groupValue,
// 		"token":    token,
// 	}
// 	return result, nil
// }

func TokenLogout(token string) error {
	duration, err := parseDuration(global.EnvConfig.Auth.Radius.TokenLifespan)
	if err != nil {
		return fmt.Errorf("error parsing duration: %v", err)
	}
	blacklistedTokens[token] = time.Now().Add(duration)
	return nil
}

func IsBlacklisted(token string) bool {
	expiry, exists := blacklistedTokens[token]
	if !exists {
		return false
	}
	return time.Now().Before(expiry)
}

// validateRadiusGroups 驗證 RADIUS 返回的 group 資訊並回傳 group value
func validateRadiusGroups(radiusOutput string) (string, error) {
	// 配置部分（根據你的實際配置結構調整）
	groupKey := global.EnvConfig.Auth.Radius.GroupKey
	allowedValues := global.EnvConfig.Auth.Radius.AllowedGroupValues

	fmt.Printf("🔍 開始解析 RADIUS 輸出，尋找 group key: %s\n", groupKey)

	lines := strings.Split(radiusOutput, "\n")
	receivedSection := false

	for i, line := range lines {
		fmt.Printf("📝 處理第 %d 行: '%s'\n", i+1, line)

		// 找到 Received Access-Accept 開始
		if strings.Contains(line, "Received Access-Accept") {
			receivedSection = true
			fmt.Printf("✅ 找到 Received Access-Accept 區段\n")
			continue
		}
		// 如果還沒進入 received section，跳過
		if !receivedSection {
			continue
		}
		// 移除行首空白
		trimmedLine := strings.TrimSpace(line)
		// 跳過空行
		if trimmedLine == "" {
			continue
		}
		// 檢查是否為屬性行（包含 = ）
		if strings.Contains(trimmedLine, "=") {
			parts := strings.SplitN(trimmedLine, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				// 移除引號
				value = strings.Trim(value, `"`)

				fmt.Printf("🔑 解析到屬性: %s = %s\n", key, value)

				// 檢查是否是目標 key
				if key == groupKey {
					fmt.Printf("🎯 找到目標 group: %s\n", value)

					// 檢查 value 是否在允許列表中
					if slices.Contains(allowedValues, value) {
						fmt.Printf("✅ Group 驗證成功，返回值: %s\n", value)
						return value, nil // 驗證成功並回傳 group value
					} else {
						return "", fmt.Errorf("group value '%s' is not in allowed list %v", value, allowedValues)
					}
				}
			}
		}
	}

	return "", fmt.Errorf("required group key '%s' not found in radius response", groupKey)
}

// // hexToString 將 hex 字串轉換為普通字串
// func hexToString(hexStr string) string {
// 	bytes, err := hex.DecodeString(hexStr)
// 	if err != nil {
// 		fmt.Printf("Error decoding hex string %s: %v\n", hexStr, err)
// 		return ""
// 	}
// 	return string(bytes)
// }

// // hexToRole 將 hex 數字對應到角色名稱
// func hexToRole(hexStr string) string {
// 	// 將 hex 轉換為整數
// 	roleNum, err := strconv.ParseInt(hexStr, 16, 64)
// 	if err != nil {
// 		fmt.Printf("Error parsing hex role %s: %v\n", hexStr, err)
// 		return "VC_USER"
// 	}

// 	// 根據數字對應角色（這些對應關係可能需要根據你的 RADIUS 伺服器調整）
// 	switch roleNum {
// 	case 0:
// 		return "VC_USER"
// 	case 1:
// 		return "VC_SUPER_USER"
// 	case 2:
// 		return "VC_ADMIN_USER"
// 	case 3:
// 		return "VC_SUPPORT_USER"
// 	default:
// 		fmt.Printf("Unknown role number: %d, using VC_USER\n", roleNum)
// 		return "VC_USER"
// 	}
// }
