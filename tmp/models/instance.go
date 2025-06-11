package models

type Instance struct {
	ID       int    `json:"id" form:"id"`
	Type     string `json:"type" form:"type"`
	Name     string `json:"name" form:"name"`
	URL      string `json:"url" form:"url"`
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
	EsUrl    string `json:"es_url" form:"es_url"`
	Auth     int    `json:"auth" form:"auth"`
}

type SSO struct {
	SsoUrl string `json:"url" form:"url"`
}
