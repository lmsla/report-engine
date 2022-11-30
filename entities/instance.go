package entities

type Instance struct {
	Common
	ID       int    `gorm:"primaryKey;index"`
	Type     string `gorm:"type:varchar(50)"`
	Name     string `gorm:"type:varchar(50)"`
	URL      string `gorm:"type:varchar(50)"`
	User     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50)"`
	Auth     bool   `type:"bool;default:false"`
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
