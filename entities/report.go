package entities

type Report struct {
	Common
	ReportID    int          `json:"report_id" form:"report_id" gorm:"primaryKey"`
	Name        string       `json:"name" form:"name"`
	Description string       `json:"description" form:"description"`
	From        string       `json:"from" form:"from"`
	To          string       `json:"to" form:"to"`
	Type        string       `json:"type" form:"type"`
	Instance    []RInstance  `json:"instances" form:"instances" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Dashboard   []RDashboard `json:"dashboards" form:"dashboards" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// type Report struct {
// 	Common
// 	ReportID    int          `json:"report_id" form:"report_id" gorm:"primaryKey"`
// 	Name        string       `json:"name" form:"name"`
// 	Description string       `json:"description" form:"description"`
// 	Type        string       `json:"type" form:"type"`
// 	Instance    []RInstance  `json:"instances" form:"instances" gorm:"many2many:InsAndDash;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// 	Dashboard   []RDashboard `json:"dashboards" form:"dashboards" gorm:"many2many:InsAndDash;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// }

// gorm:"many2many:dashboards_roles;joinForeignKey:DashboardID;joinReferences:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;

type RInstance struct {
	Common1
	ReportID     int    `json:"report_id" form:"report_id"`
	InstanceName string `json:"instance_name" form:"instance_name"`
	// Dashboard []RDashboard `gorm:"many2many:ins_and_dash;"`
}

type RDashboard struct {
	Common1
	ReportID int    `json:"report_id" form:"report_id"`
	UID      string `json:"uid" form:"uid" gorm:"type:varchar(50)"`
	// Instance []RInstance `gorm:"many2many:ins_and_dash;"`
	DashboardName string `json:"dashboard_name" form:"dashboard_name"`
}

// Dashboards  []Dashboard `json:"dashboards" form:"dashboards" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
