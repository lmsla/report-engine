package entities

type Instance struct {
	Common
	ID       int    `gorm:"primaryKey;index" json:"id" form:"id"`
	Type     string `gorm:"type:varchar(50)" json:"type" form:"type"`
	Name     string `gorm:"type:varchar(50)" json:"name" form:"name"`
	URL      string `gorm:"type:varchar(50)" json:"url" form:"url"`
	User     string `gorm:"type:varchar(50)" json:"user" form:"user"`
	Password string `gorm:"type:varchar(50)" json:"password" form:"password"`
	EsUrl    string `gorm:"type:varchar(50)" json:"es_url" form:"es_url"`
	Auth     int   `type:"int;default:0" json:"auth" form:"auth"`
}

type Dashboard struct {
	Common
	DashboardID int    `json:"dashboard_id" form:"dashboard_id" gorm:"primaryKey"`
	Name        string `json:"dashboard_name" form:"dashboard_name"`
	InstanceID  int    `json:"instance_id" form:"instance_id"`
	UID         string `json:"uid" form:"uid" gorm:"type:varchar(50)"`
	// Reports       []Report `gorm:"many2many:Report_Dashboards;"`
}

type Visualization struct {
	Common
	DashboardID int    `json:"dashboard_id" form:"dashboard_id" gorm:"primaryKey"`
	Name        string `json:"dashboard_name" form:"dashboard_name"`
	InstanceID  int    `json:"instance_id" form:"instance_id"`
	UID         string `json:"uid" form:"uid" gorm:"type:varchar(50)"`
	// Reports       []Report `gorm:"many2many:Report_Dashboards;"`
}
