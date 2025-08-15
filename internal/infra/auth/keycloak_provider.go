package auth

import (
	"context"
	"crypto/tls"
	"fmt"

	"report-backend-golang/global"

	"github.com/Nerzal/gocloak/v13"
)

type KeycloakProvider struct {
	client *gocloak.GoCloak
	realm  string
}

// NewKeycloakProvider 建立 Keycloak 認證提供者
func NewKeycloakProvider() *KeycloakProvider {
	url := global.EnvConfig.Auth.Keycloak.Url
	client := gocloak.NewClient(url)
	restyClient := client.RestyClient()
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	return &KeycloakProvider{
		client: client,
		realm:  global.EnvConfig.Auth.Keycloak.Realm,
	}
}

// Authenticate 驗證用戶憑證
func (p *KeycloakProvider) Authenticate(ctx context.Context, creds map[string]string) (map[string]interface{}, error) {
	// fmt.Printf("Keycloak Authenticate called with creds: %+v\n", creds)

	var token string
	var err error

	// 檢查是否有 token
	if tokenValue, ok := creds["token"]; ok {
		token = tokenValue
	} else if username, ok := creds["username"]; ok {
		// 使用 username/password 登入取得 token
		password, exists := creds["password"]
		if !exists {
			return nil, fmt.Errorf("password is required when username is provided")
		}

		// 嘗試用提供的用戶名密碼登入
		userToken, err := p.client.Login(ctx, global.EnvConfig.SSO.ClientID, "", p.realm, username, password)

		if err != nil {
			return nil, fmt.Errorf("keycloak login failed: %w", err)
		}

		token = userToken.AccessToken
	} else {
		return nil, fmt.Errorf("token or username/password is required")
	}

	// 透過 token 取得用戶資訊
	userInfo, err := p.client.GetUserInfo(ctx, token, p.realm)
	if err != nil {
		return nil, fmt.Errorf("keycloak authentication failed: %w", err)
	}

	// 將 userInfo 轉換為 map 格式，包含 token
	result := map[string]interface{}{
		"sub":                *userInfo.Sub,
		"preferred_username": *userInfo.PreferredUsername,
		"realm":              p.realm,
		"source":             "keycloak",
		"access_token":       token,
	}

	// fmt.Printf("Keycloak result: %+v\n", result)
	return result, nil
}

// Profile 從認證結果建立用戶資料
func (p *KeycloakProvider) Profile(raw map[string]interface{}) (*SSOUser, error) {
	sub, ok := raw["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid user id")
	}

	username, ok := raw["preferred_username"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid username")
	}

	realm, _ := raw["realm"].(string)

	user := &SSOUser{
		ID:     sub,
		Name:   username,
		Domain: realm,
		Roles:  []string{}, // 可根據需要從 Keycloak 取得角色資訊
	}

	return user, nil
}
 