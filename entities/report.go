package entities

type Report struct {
	Common
	ReportID    int         `json:"report_id" form:"report_id" gorm:"primaryKey"`
	Name        string      `json:"name" form:"name"`
	Description string      `json:"description" form:"description"`
	From        string      `json:"from" form:"from"`
	To          string      `json:"to" form:"to"`
	Type        string      `json:"type" form:"type"`
	// Instance    []RInstance  `json:"Instance" form:"Instance" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Dashboard   []RDashboard `json:"Dashboard" form:"Dashboard" gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Instance    []RInstance  `gorm:"many2many:ReportInstance;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Dashboard   []RDashboard  `gorm:"many2many:ReportDashboard;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Instance    []RInstance  
	Dashboard   []RDashboard
}

// joinForeignKey:RoleID;joinReferences:UserID


// `json:"report" form:"report" gorm:"-"`
type Report1 struct {
	Common
	ReportID    int          `json:"report_id" form:"report_id" gorm:"primaryKey"`
	Name        string       `json:"name" form:"name"`
	Description string       `json:"description" form:"description"`
	From        string       `json:"from" form:"from"`
	To          string       `json:"to" form:"to"`
	Type        string       `json:"type" form:"type"`
	Instance    []RInstance  `gorm:"many2many:ReportInstance;joinForeignKey:ReportID;joinReferences:InstanceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Dashboard   []RDashboard `json:"dashboards" form:"dashboards" gorm:"many2many:ReportDashboard;joinForeignKey:ReportID;joinReferences:DashboardID;joinReferences:InstanceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ReportDashboard struct {
	ReportID    int
	DashboardID int
	InstanceID int
	// DashboardName string `json:"dashboard_name" form:"dashboard_name"`
}

type ReportInstance struct {
	ReportID   int
	InstanceID int
	// InstanceName string `json:"name" form:"name"`
}

type RInstance struct {
	Common
	InstanceID   int          `json:"instance_id" form:"instance_id" gorm:"primaryKey"`
	ReportID     int          `json:"report_id" form:"report_id"`
	InstanceName string       `json:"instance_name" form:"instance_name"`
	IP           string       `json:"ip" form:"ip"`
	Type         string       `json:"type" form:"type"`
	User         string      `json:"user" form:"user"`
	Pass         string      `json:"pass" form:"pass"`
	Dashboard    []RDashboard `gorm:"foreignKey:InstanceID"`
}

type RDashboard struct {
	Common
	DashboardID int    `json:"dashboard_id" form:"dashboard_id" gorm:"primaryKey"`
	InstanceID  int    `json:"instance_id" form:"instance_id"`
	ReportID    int    `json:"report_id" form:"report_id"`
	UID         string `json:"uid" form:"uid" gorm:"type:varchar(50)"`
	// Instance []RInstance `gorm:"many2many:ins_and_dash;"`
	DashboardName string `json:"dashboard_name" form:"dashboard_name"`
}
