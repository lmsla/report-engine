package entities

type User struct {
	Common
	UserID   int    `json:"user_id" form:"user_id" gorm:"primaryKey"`
	UserName string `json:"user_name" form:"user_name"`
	Roles    []Role `json:"roles" form:"roles" gorm:"many2many:users_roles;joinForeignKey:UserID;joinReferences:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Role struct {
	Common
	RoleID     int         `json:"role_id" form:"role_id" gorm:"primaryKey"`
	RoleName   string      `json:"role_name" form:"role_name"`
	Dashboards []Dashboard `json:"dashboards" form:"dashboards" gorm:"many2many:dashboards_roles;joinForeignKey:RoleID;joinReferences:DashboardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Users      []User      `gorm:"many2many:users_roles;joinForeignKey:RoleID;joinReferences:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// users_roles
type UsersRole struct {
	UserID int
	RoleID int
}
