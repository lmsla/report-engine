# 認證機制擴充技術規範文件

此文件說明如何在現有 Keycloak 驗證基礎上，新增 RADIUS 驗證並實現動態切換，供 Cursor 自動實作。

## 一、配置更新

```yaml
# config/setting.yml
auth:
  type: "keycloak"   # or "radius"
  keycloak:
    url: "https://sso.example.com"
    realm: "master"
    user: "svc-admin"
    secret: "svc-secret"
  radius:
    server: "10.99.1.248"
    secret: "radius123"
    nas_port: "0"
    timeout_seconds: 3

說明：
	•	auth.type 決定啟動使用的 Provider。
	•	支援從環境變數覆蓋：AUTH_TYPE, RADIUS_SERVER, RADIUS_SECRET 等。

二、通用介面定義

// internal/infra/auth/provider.go
package auth

import "context"

type SSOUser struct {
  ID     string
  Name   string
  Domain string   
  Roles  []string 
}

type AuthProvider interface {
  Authenticate(ctx context.Context, creds map[string]string) (map[string]interface{}, error)
  Profile(raw map[string]interface{}) (*SSOUser, error)
}

三、Keycloak Provider 實作

// internal/infra/auth/keycloak_provider.go
package auth

import (
  "context"
  "crypto/tls"
  "github.com/Nerzal/gocloak/v13"
  "report-backend-golang/global"
)

type KeycloakProvider struct { client gocloak.GoCloak; realm string }

func NewKeycloakProvider() *KeycloakProvider { /* ... */ }
func (p *KeycloakProvider) Authenticate(ctx context.Context, creds map[string]string) (map[string]interface{}, error) { /* ... */ }
func (p *KeycloakProvider) Profile(raw map[string]interface{}) (*SSOUser, error) { /* ... */ }

四、Radius Provider 實作

// internal/infra/auth/radius_provider.go
package auth

import (
  "context"; "fmt"; "net"; "strconv"; "time"
  "github.com/bronze1man/radius"
  "report-backend-golang/global"
)

type RadiusProvider struct { addr, secret string; timeout time.Duration }
func NewRadiusProvider() *RadiusProvider { /* ... */ }
func (r *RadiusProvider) Authenticate(ctx context.Context, creds map[string]string) (map[string]interface{}, error) { /* ... */ }
func (r *RadiusProvider) Profile(raw map[string]interface{}) (*SSOUser, error) { /* ... */ }

五、Provider Factory

// cmd/server/main.go
func initAuthProvider() auth.AuthProvider {
  switch global.EnvConfig.Auth.Type {
  case "radius": return auth.NewRadiusProvider()
  default:       return auth.NewKeycloakProvider()
  }
}

六、Middleware 改造

func AuthMiddleware(p auth.AuthProvider) gin.HandlerFunc {
  return func(c *gin.Context) {
    creds := map[string]string{ /* username, password, nas_port */ }
    raw, err := p.Authenticate(c, creds)
    // error handling...
    user, err := p.Profile(raw)
    c.Set("user", user)
    c.Next()
  }
}

七、路由與 Handler
	•	/login: 使用上述 Middleware 完成驗證、發行 JWT
	•	/userinfo: 回傳 SSOUser 資訊

八、日誌追蹤與錯誤處理
	•	在各 Provider Authenticate 與 Profile 中 log 事件（不含敏感資訊）
	•	統一回應格式：
	•	401 Unauthorized
	•	403 Forbidden
	•	500 Internal Server Error

