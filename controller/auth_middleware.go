package controller

import (
	"encoding/base64"
	// "fmt"
	"net/http"
	"strings"

	"report-backend-golang/global"
	"report-backend-golang/internal/infra/auth"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)

var authProvider auth.AuthProvider

// InitAuthProvider 初始化認證提供者
func InitAuthProvider() {
	authProvider = auth.InitAuthProvider()
}

// GetAuthProvider 取得已初始化的認證提供者
func GetAuthProvider() auth.AuthProvider {
	return authProvider
}

// AuthMiddleware 統一認證中間件，支援 Bearer Token 和 Basic Auth
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			return
		}

		var creds map[string]string

		if strings.HasPrefix(authHeader, "Basic ") {
			// Basic Auth for RADIUS
			basicAuth := strings.TrimPrefix(authHeader, "Basic ")
			decoded, err := base64.StdEncoding.DecodeString(basicAuth)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid basic auth format",
				})
				return
			}

			parts := strings.SplitN(string(decoded), ":", 2)
			if len(parts) != 2 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "invalid basic auth format",
				})
				return
			}

			creds = map[string]string{
				"username": parts[0],
				"password": parts[1],
			}
		} else if strings.HasPrefix(authHeader, "Bearer ") {
			// Bearer Token for Keycloak or RADIUS JWT
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// 檢查是否為 RADIUS 認證類型
			if global.EnvConfig.Auth.Type == "radius" {
				// 使用 JWT Token 驗證
				userInfo, err := services.ValidateTokenAndGetUserInfo(token)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "invalid token: " + err.Error(),
					})
					return
				}

				// 建立用戶資料
				user := &auth.SSOUser{
					ID:     userInfo.UserName,
					Name:   userInfo.UserName,
					Domain: "radius",
					Group:  userInfo.UserGroup,
				}

				c.Set("user", user)
				c.Next()
				return
			} else {
				// Keycloak Token 驗證
				creds = map[string]string{
					"token": token,
				}
			}
		} else {
			// 處理沒有 Bearer 前綴的 token（非標準格式，僅為兼容性）
			// 警告：這違反了 HTTP 標準，應該修改前端使用 Bearer 前綴
			token := authHeader

			// 記錄警告日誌
			// fmt.Printf("WARNING: 收到非標準的 Authorization header，缺少 Bearer 前綴: %s\n", authHeader[:min(20, len(authHeader))])

			// 檢查是否為 RADIUS 認證類型
			if global.EnvConfig.Auth.Type == "radius" {
				// 使用 JWT Token 驗證
				userInfo, err := services.ValidateTokenAndGetUserInfo(token)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "invalid token: " + err.Error(),
					})
					return
				}

				// 建立用戶資料
				user := &auth.SSOUser{
					ID:     userInfo.UserName,
					Name:   userInfo.UserName,
					Domain: "radius",
					Group:  userInfo.UserGroup,
				}

				c.Set("user", user)
				c.Next()
				return
			} else {
				// Keycloak Token 驗證
				creds = map[string]string{
					"token": token,
				}
			}
		}

		// 執行認證
		raw, err := authProvider.Authenticate(c.Request.Context(), creds)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authentication failed: " + err.Error(),
			})
			return
		}

		// 取得用戶資料
		user, err := authProvider.Profile(raw)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get user profile: " + err.Error(),
			})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// GetUserInfo 相容性函數，保持向後相容
func GetUserInfo(c *gin.Context) {
	AuthMiddleware()(c)
}
