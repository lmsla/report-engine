package models

type Response struct {
	Msg     string
	Success bool
}

type Dashboard struct {
	DashboardID   int    `json:"dashboard_id" form:"dashboard_id"`
	DashboardName string `json:"dashboard_name" form:"dashboard_name"`
	Alias         string `json:"alias" form:"alias"`
	Description   string `json:"description" form:"description"`
	UID           string `json:"uid" form:"uid"`
	InstanceID    int    `json:"instance_id" form:"instance_id"`
	InstanceName  string `json:"instance_name" form:"instance_name"`
	MenuID        int    `json:"menu_id" form:"menu_id"`
	MenuName      string `json:"menu_name" form:"menu_name"`
}

type Role struct {
	RoleID     int         `json:"role_id" form:"role_id"`
	RoleName   string      `json:"role_name" form:"role_name"`
	Dashboards []Dashboard `json:"dashboards" form:"dashboards"`
}

type Menu struct {
	MenuID      int         `json:"menu_id" form:"menu_id"`
	MenuName    string      `json:"menu_name" form:"menu_name"`
	Sort        int         `json:"sort" form:"sort"`
	Dashoboards []Dashboard `json:"dashboards" form:"dashboards"`
}
type UpdateDashboard struct {
	OldID int `json:"old_dashboard_id" form:"old_dashboard_id"`
	NewID int `json:"new_dashboard_id" form:"new_dashboard_id"`
}

type DashboardMenu struct {
	Children    []DashboardMenu  `json:"children"`
	ID          int              `json:"id"`
	UID			string           `json:"uid"`
	RouterPath  string           `json:"router_path"`
	Title       string           `json:"title"`
	Sort        int              `json:"sort"`
	ParentID    int              `json:"parent_id"`
	Type        string           `json:"type"`
}
