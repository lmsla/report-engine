package models

type Response struct {
	Success bool   `json:"success" form:"success"`
	Msg     string `json:"msg" form:"msg"`
	Body    interface{}
}

// ErrorResponse 錯誤回應結構
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// RadiusLoginRequest RADIUS 登入請求
type RadiusLoginRequest struct {
	UserName     string `json:"user_name" binding:"required" example:"admin"`
	UserPassword string `json:"user_password" binding:"required" example:"password"`
}

// RadiusLoginResponse RADIUS 登入回應
type RadiusLoginResponse struct {
	AccessAccept bool   `json:"access_accept" example:"true"`
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	UserDomain   string `json:"user_domain" example:"OPERATOR"`
	UserRole     string `json:"user_role" example:"VC_SUPER_USER"`
}

// LogoutRequest 登出請求
type LogoutRequest struct {
	AccessToken string `json:"access_token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	AccessAccept bool   `json:"access_accept" binding:"required" example:"true"`
}

// LogoutResponse 登出回應
type LogoutResponse struct {
	Success bool   `json:"success" example:"true"`
	Msg     string `json:"msg" example:"user logout success"`
}

// LoginRequest 統一登入請求
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"password"`
	NasPort  string `json:"nas_port,omitempty" example:"0"`
}

// UserInfo 用戶資訊
type UserInfo struct {
	ID     string   `json:"id" example:"admin"`
	Name   string   `json:"name" example:"admin"`
	Domain string   `json:"domain" example:"radius"`
	Roles  []string `json:"roles" example:"VC_SUPER_USER"`
}

// LoginResponse 統一登入回應
type LoginResponse struct {
	Token string   `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserInfo `json:"user"`
}
