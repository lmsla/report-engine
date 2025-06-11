package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 統一登入端點 (適用於 RADIUS)
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
