package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)

// Login 統一登入端點 (適用於 RADIUS)
// @Summary 統一登入端點
// @Description 支援 RADIUS 和 Keycloak 認證的統一登入端點
// @Tags Authentication
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "登入資訊"
// @Success 200 {object} models.LoginResponse "登入成功"
// @Failure 400 {object} models.ErrorResponse "請求格式錯誤"
// @Failure 401 {object} models.ErrorResponse "認證失敗"
// @Failure 500 {object} models.ErrorResponse "內部錯誤"
// @Router /login [post]
func Login(c *gin.Context) {
	fmt.Println("Login function called")

	var loginReq struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		NasPort  string `json:"nas_port,omitempty"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request: " + err.Error(),
		})
		return
	}

	// 準備認證憑證
	creds := map[string]string{
		"username": loginReq.Username,
		"password": loginReq.Password,
	}
	if loginReq.NasPort != "" {
		creds["nas_port"] = loginReq.NasPort
	}

	// 執行認證 - 使用已初始化的 authProvider
	provider := GetAuthProvider()
	if provider == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "auth provider not initialized",
		})
		return
	}

	fmt.Printf("About to call provider.Authenticate with creds: %+v\n", creds)
	raw, err := provider.Authenticate(c.Request.Context(), creds)
	fmt.Printf("provider.Authenticate returned: raw=%+v, err=%v\n", raw, err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication failed: " + err.Error(),
		})
		return
	}

	// 取得用戶資料
	user, err := provider.Profile(raw)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user profile: " + err.Error(),
		})
		return
	}

	// 準備回應
	response := gin.H{
		"user": user,
	}

	// 如果認證結果中有 access_token，則使用它（用於 Keycloak）
	if accessToken, ok := raw["access_token"].(string); ok {
		response["token"] = accessToken
	} else {
		// 對於 RADIUS，我們可以產生一個簡單的 token (這裡只是示範)
		// 在實際應用中，你可能想要使用 JWT 或其他方式
		token := base64.StdEncoding.EncodeToString([]byte(user.ID + ":" + user.Name))
		response["token"] = token
	}

	c.JSON(http.StatusOK, response)
}

// GetAccessToken RADIUS 登入取得 Token
// @Summary RADIUS 登入取得 Token
// @Description 使用 RADIUS 認證取得 JWT Token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param account body models.RadiusLoginRequest true "RADIUS 帳戶資訊"
// @Success 200 {object} models.RadiusLoginResponse "認證成功"
// @Failure 400 {object} models.ErrorResponse "請求格式錯誤"
// @Failure 401 {object} models.ErrorResponse "認證失敗"
// @Router /token [post]
func GetAccessToken(c *gin.Context) {
	var loginReq struct {
		UserName     string `json:"user_name" binding:"required"`
		UserPassword string `json:"user_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request: " + err.Error(),
		})
		return
	}

	// 使用 RADIUS 認證服務
	result, err := services.RadiusAuthenticate(loginReq.UserName, loginReq.UserPassword)
	if err != nil {
		// 認證失敗，返回統一格式但保持 401 status
		response := gin.H{
			"access_accept": false,
			"access_token":  "",
			"user_domain":   "",
			"user_role":     "",
		}
		c.JSON(http.StatusUnauthorized, response) // 保持 401
		return
	}
	fmt.Println("result", result)
	// 準備回應
	response := gin.H{
		"access_accept": true,
		"access_token":  result["token"],
		"user_domain":   "OPERATOR",
		"user_role":     result["role"],
		"user_group":    result["group"],
	}

	c.JSON(http.StatusOK, response)
}

// TokenLogout RADIUS Token 登出
// @Summary RADIUS Token 登出
// @Description 登出並使 Token 失效
// @Tags Authentication
// @Accept json
// @Produce json
// @Param token body models.LogoutRequest true "Token 資訊"
// @Success 200 {object} models.LogoutResponse "登出成功"
// @Failure 400 {object} models.ErrorResponse "請求格式錯誤"
// @Failure 500 {object} models.ErrorResponse "登出失敗"
// @Router /logout [post]
func TokenLogout(c *gin.Context) {
	var logoutReq struct {
		AccessAccept bool   `json:"access_accept" binding:"required"`
		AccessToken string `json:"access_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&logoutReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request: " + err.Error(),
		})
		return
	}

	// 執行登出
	err := services.TokenLogout(logoutReq.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "logout failed: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "user logout success",
	})
}

// UserInfo 取得當前用戶資訊
func UserInfo(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found in context",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
