package models

type Instance struct {
	ID       int    `json:"id" form:"id"`
	Type     string `json:"type" form:"type"`
	Name     string `json:"name" form:"name"`
	URL      string `json:"url" form:"url"`
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
	Auth     bool   `json:"auth" form:"auth"`
}
