package entities

// import "time"

type Instance struct {
	Common
	InstanceID   int         `json:"instance_id" form:"instance_id" gorm:"primaryKey"`
	IP           string      `json:"ip" form:"ip"`
	InstanceName string      `json:"name" form:"name"`
	Type         string      `json:"type" form:"type"`
	User         string      `json:"user" form:"user"`
	Pass         string      `json:"pass" form:"pass"`
	Auth         bool        `json:"auth" form:"auth"`
	Dashboards   []Dashboard `gorm:"foreignKey:InstanceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Dashboard struct {
	Common
	DashboardID   int    `json:"dashboard_id" form:"dashboard_id" gorm:"primaryKey"`
	DashboardName string `json:"dashboard_name" form:"dashboard_name"`
	Alias         string `json:"alias" form:"alias"`
	Description   string `json:"description" form:"description"`
	UID           string `json:"uid" form:"uid" gorm:"type:varchar(50)"`
	InstanceID    int    `json:"instance_id" form:"instance_id"`
	MenuID        int    `json:"menu_id" form:"menu_id"`
	Roles         []Role `json:"roles" form:"roles" gorm:"many2many:dashboards_roles;joinForeignKey:DashboardID;joinReferences:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//Menus         []Menu `gorm:"many2many:menus_dashboards;joinForeignKey:DashboardID;joinReferences:MenuID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Menu struct {
	MenuID     int         `json:"menu_id" form:"menu_id" gorm:"primaryKey"`
	MenuName   string      `json:"menu_name" form:"menu_name"`
	Sort       int         `json:"sort" form:"sort"`
	Dashboards []Dashboard `json:"dashboards" form:"dashboards" gorm:"foreignKey:MenuID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Group struct {
	Common
	GroupID     int           `json:"group_id" form:"group_id" gorm:"primaryKey"`
	Name        string        `json:"group_name" form:"group_name"`
	GroupMember []GroupMember `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type GroupMember struct {
	Common
	MemberID   int    `json:"member_id" form:"member_id" gorm:"primaryKey"`
	GroupID    int    `json:"group_id" form:"group_id"`
	MemberName string `json:"member_name" form:"member_name"`
	Email      string `json:"email" form:"email"`
}
