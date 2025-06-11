package auth

import (
	"context"
	"fmt"
	"time"

	"report-backend-golang/global"
)

type RadiusProvider struct {
	addr    string
	secret  string
	timeout time.Duration
}

// NewRadiusProvider 建立 RADIUS 認證提供者
func NewRadiusProvider() *RadiusProvider {
	// 設定預設值
	server := "127.0.0.1:1812"
	secret := "testing123"
	timeout := 3

	// 如果有配置則使用配置值
	if global.EnvConfig.Auth.Radius.Server != "" {
		server = global.EnvConfig.Auth.Radius.Server
	}
	if global.EnvConfig.Auth.Radius.Secret != "" {
		secret = global.EnvConfig.Auth.Radius.Secret
	}
	if global.EnvConfig.Auth.Radius.TimeoutSeconds > 0 {
		timeout = global.EnvConfig.Auth.Radius.TimeoutSeconds
	}

	return &RadiusProvider{
		addr:    server,
		secret:  secret,
		timeout: time.Duration(timeout) * time.Second,
	}
}

// Authenticate 驗證用戶憑證
func (r *RadiusProvider) Authenticate(ctx context.Context, creds map[string]string) (map[string]interface{}, error) {
	username, ok := creds["username"]
	if !ok {
		return nil, fmt.Errorf("username is required")
	}

	password, ok := creds["password"]
	if !ok {
		return nil, fmt.Errorf("password is required")
	}

	// 簡化版本：先跳過真正的 RADIUS 連接，用於測試流程
	// TODO: 實作真正的 RADIUS 認證
	fmt.Printf("RADIUS 認證測試 - 用戶: %s, 密碼: %s, 伺服器: %s\n", username, password, r.addr)

	// 模擬認證成功
	if username != "" && password != "" {
		// 建立認證結果
		result := map[string]interface{}{
			"username": username,
			"source":   "radius",
			"server":   r.addr,
		}
		return result, nil
	}

	return nil, fmt.Errorf("radius authentication failed: invalid credentials")
}

// Profile 從認證結果建立用戶資料
func (r *RadiusProvider) Profile(raw map[string]interface{}) (*SSOUser, error) {
	username, ok := raw["username"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid username")
	}

	user := &SSOUser{
		ID:     username, // RADIUS 通常以用戶名作為 ID
		Name:   username,
		Domain: "radius",
		Roles:  []string{}, // RADIUS 本身不提供角色資訊，可根據需要擴展
	}

	return user, nil
}
