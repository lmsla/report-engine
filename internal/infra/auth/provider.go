package auth

import "context"

// SSOUser 統一的用戶結構
type SSOUser struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Domain string   `json:"domain"`
	Roles  []string `json:"roles"`
}

// AuthProvider 認證提供者介面
type AuthProvider interface {
	Authenticate(ctx context.Context, creds map[string]string) (map[string]interface{}, error)
	Profile(raw map[string]interface{}) (*SSOUser, error)
}
