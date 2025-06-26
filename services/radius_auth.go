package services

import (
	"bytes"
	// "encoding/hex"
	"fmt"
	"os/exec"

	// "regexp"
	"report-backend-golang/global"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
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

// RadiusAuthenticate 真實的 RADIUS 認證函數
func RadiusAuthenticate(username, password string) (map[string]interface{}, error) {
	// 取得 RADIUS 配置
	server := global.EnvConfig.Auth.Radius.Server
	nasPort := global.EnvConfig.Auth.Radius.NasPort
	secret := global.EnvConfig.Auth.Radius.Secret
	radtestPath := global.EnvConfig.Auth.Radius.RadtestPath

	if server == "" || nasPort == "" || secret == "" {
		return nil, fmt.Errorf("radius configuration incomplete")
	}

	fmt.Printf("RADIUS 真實認證 - 用戶: %s, 伺服器: %s:%s\n", username, server, nasPort)

	// 檢查 radtest 路徑
	if radtestPath == "" {
		radtestPath = "radtest" // 預設值，依賴 PATH
	}

	// 執行 radtest 命令
	cmd := exec.Command(radtestPath, username, password, server, nasPort, secret)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("RADIUS 認證失敗: %v\nStderr: %s\n", err, stderr.String())
		return nil, fmt.Errorf("radius authentication failed: invalid credentials")
	}

	output := out.String()
	fmt.Printf("RADIUS 回應: %s\n", output)

	// 檢查是否認證成功
	if !strings.Contains(output, "Access-Accept") {
		return nil, fmt.Errorf("radius authentication failed: access denied")
	}

	// var userDomain, userRole string

	// // 方法1：嘗試文字格式解析（rule-engine 相容格式）
	// fmt.Println("嘗試文字格式解析...")
	// reDomainText := regexp.MustCompile(`VC_USER_DOMAIN\s*=\s*"([^"]+)"`)
	// reRoleText := regexp.MustCompile(`VC_USER_ROLE\s*=\s*([^\s]+)`)
	// domainTextMatch := reDomainText.FindStringSubmatch(output)
	// roleTextMatch := reRoleText.FindStringSubmatch(output)

	// if len(domainTextMatch) >= 2 && len(roleTextMatch) >= 2 {
	// 	// 找到文字格式，使用它
	// 	userDomain = domainTextMatch[1]
	// 	userRole = roleTextMatch[1]
	// 	fmt.Printf("✅ 使用文字格式解析 - Domain: %s, Role: %s\n", userDomain, userRole)
	// } else {
	// 	// 方法2：嘗試 VSA 數字格式解析
	// 	fmt.Println("文字格式未找到，嘗試 VSA 格式解析...")
	// 	reVSA1 := regexp.MustCompile(`Attr-26\.45346\.1\s*=\s*0x([0-9a-fA-F]+)`) // Domain
	// 	reVSA2 := regexp.MustCompile(`Attr-26\.45346\.2\s*=\s*0x([0-9a-fA-F]+)`) // Role
	// 	vsa1Match := reVSA1.FindStringSubmatch(output)
	// 	vsa2Match := reVSA2.FindStringSubmatch(output)

	// 	// 解析 Domain (VSA1)
	// 	if len(vsa1Match) >= 2 {
	// 		userDomain = hexToString(vsa1Match[1])
	// 		fmt.Printf("✅ VSA Domain 解析: %s (來自 hex: %s)\n", userDomain, vsa1Match[1])
	// 	} else {
	// 		fmt.Println("⚠️  未找到 VSA Domain 屬性，使用預設值")
	// 		userDomain = "OPERATOR"
	// 	}

	// 	// 解析 Role (VSA2)
	// 	if len(vsa2Match) >= 2 {
	// 		userRole = hexToRole(vsa2Match[1])
	// 		fmt.Printf("✅ VSA Role 解析: %s (來自 hex: %s)\n", userRole, vsa2Match[1])
	// 	} else {
	// 		fmt.Println("⚠️  未找到 VSA Role 屬性，使用預設值")
	// 		userRole = "VC_USER"
	// 	}

	// 	if len(vsa1Match) >= 2 || len(vsa2Match) >= 2 {
	// 		fmt.Println("✅ 使用 VSA 格式解析")
	// 	} else {
	// 		fmt.Println("⚠️  兩種格式都未找到，使用預設值")
	// 	}
	// }

	// // 確保有預設值
	// if userDomain == "" {
	// 	userDomain = "OPERATOR"
	// 	fmt.Println("🔧 Domain 為空，使用預設值: OPERATOR")
	// }
	// if userRole == "" {
	// 	userRole = "VC_USER"
	// 	fmt.Println("🔧 Role 為空，使用預設值: VC_USER")
	// }

	// fmt.Printf("📋 最終解析結果 - Domain: %s, Role: %s\n", userDomain, userRole)

	// // 檢查角色權限
	// if len(global.EnvConfig.Auth.Radius.AdminVcRole) > 0 {
	// 	if !slices.Contains(global.EnvConfig.Auth.Radius.AdminVcRole, userRole) {
	// 		return nil, fmt.Errorf("user role not authorized: %s", userRole)
	// 	}
	// }

	// 新增 group 驗證
	var groupValue string

	groupValue, err = validateRadiusGroups(output)
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
		"source":   "radius",
		"server":   server,
		"group":    groupValue,
		"token":    token,
	}
	return result, nil
}

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
