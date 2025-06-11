package auth

import (
	"fmt"
	"report-backend-golang/global"
)

// InitAuthProvider 根據配置初始化認證提供者
func InitAuthProvider() AuthProvider {
	authType := "keycloak" // 預設使用 keycloak

	// 從配置中獲取認證類型
	if global.EnvConfig.Auth.Type != "" {
		authType = global.EnvConfig.Auth.Type
	}

	// Debug 輸出
	fmt.Printf("Auth Type: %s\n", authType)
	fmt.Printf("Radius Server: %s\n", global.EnvConfig.Auth.Radius.Server)

	switch authType {
	case "radius":
		return NewRadiusProvider()
	case "keycloak":
		fallthrough
	default:
		return NewKeycloakProvider()
	}
}
