package entities

import (
	"os"
	"report-backend-golang/global"

	"github.com/gookit/color"
)

func InitTable() {

	var err error

	err = global.Mysql.Migrator().DropTable(
		User{},
		// Role{},
		Instance{},
		Dashboard{},
		// Menu{},
		Group{},
		GroupMember{},
		Report{},
		Schedule{},
		ReportDashboard{},
		ReportInstance{},
		RInstance{},
		RDashboard{},
		FileHistory{},
		"users_roles",
		"dashboards_roles",
		"menus_dashboards",
	)
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化數據失敗(移除原始Tables),err: %v\n", err)
		os.Exit(0)
	}

	err = global.Mysql.AutoMigrate(
		User{},
		// Role{},
		Instance{},
		Dashboard{},
		// Menu{},
		Group{},
		GroupMember{},
		Report{},
		Schedule{},
		ReportDashboard{},
		ReportInstance{},
		RInstance{},
		RDashboard{},
		FileHistory{},
	)

	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化數據表失敗(建立Tables),err: %v\n", err)
		os.Exit(0)
	}

	// err = global.Mysql.Omit("Roles.Dashboards").Create(&UserData).Error
	// if err != nil {
	// 	color.Warn.Printf("[Mysql]-->初始化User數據表失敗,err: %v\n", err)
	// 	os.Exit(0)
	// }

	// err = global.Mysql.Create(&InstanceData).Error
	// if err != nil {
	// 	color.Warn.Printf("[Mysql]-->初始化Inventory數據表失敗,err: %v\n", err)
	// 	os.Exit(0)
	// }

	// // 寫入Relation Table (dashboards_roles)
	// for i := 0; i < len(UserData); i++ {
	// 	for j := 0; j < len(UserData[i].Roles); j++ {
	// 		for k := 0; k < len(UserData[i].Roles[j].Dashboards); k++ {
	// 			var dashboardID, roleID int
	// 			global.Mysql.Table("dashboards").Where("dashboard_name = ?", UserData[i].Roles[j].Dashboards[k].DashboardName).Select("dashboard_id").Scan(&dashboardID)
	// 			global.Mysql.Table("roles").Where("role_name = ?", UserData[i].Roles[j].RoleName).Select("role_id").Scan(&roleID)
	// 			err := global.Mysql.Table("dashboards_roles").Create(map[string]interface{}{"dashboard_id": dashboardID, "role_id": roleID}).Error
	// 			if err != nil {
	// 				color.Warn.Printf("[Mysql]-->初始化dashboard_roles數據表失敗,err: %v\n", err)
	// 				os.Exit(0)
	// 			}
	// 		}

	// 	}
	// }

	// err = global.Mysql.Omit("Dashboards").Create(&MenuData).Error
	// if err != nil {
	// 	color.Warn.Printf("[Mysql]-->初始化Menu數據表失敗,err: %v\n", err)
	// 	os.Exit(0)
	// }

	// for i := 0; i < len(MenuData.Dashboards); i++ {
	// 	var menuID int
	// 	global.Mysql.Table("menus").Where("menu_name = ?", MenuData.MenuName).Select("menu_id").Scan(&menuID)
	// 	//global.Mysql.Table("dashboards").Where("dashboard_name = ?", MenuData.Dashboards[i].DashboardName).Select("dashboard_id").Scan(&dashboardID)
	// 	//err := global.Mysql.Table("menus_dashboards").Create(map[string]interface{}{"dashboard_id": dashboardID, "menu_id": menuID}).Error
	// 	err := global.Mysql.Table("dashboards").Where("dashboard_name = ?", MenuData.Dashboards[i].DashboardName).Update("menu_id", menuID).Error
	// 	if err != nil {
	// 		color.Warn.Printf("[Mysql]-->初始化menus_dashboards數據表失敗,err: %v\n", err)
	// 		os.Exit(0)
	// 	}
	// }

	global.Mysql.Migrator().CreateConstraint(&Instance{}, "Dashboards")
	global.Mysql.Migrator().CreateConstraint(&User{}, "Roles")
	global.Mysql.Migrator().CreateConstraint(&Role{}, "Users")
	global.Mysql.Migrator().CreateConstraint(&Role{}, "Dashboards")
	// global.Mysql.Migrator().CreateConstraint(&Dashboard{}, "Roles")
	// global.Mysql.Migrator().CreateConstraint(&Menu{}, "Dashboards")
	global.Mysql.Migrator().CreateConstraint(&Group{}, "Group")
	global.Mysql.Migrator().CreateConstraint(&GroupMember{}, "GroupMember")
	//global.Mysql.Migrator().CreateConstraint(&Dashboard{}, "Menus")

	color.Info.Println("[Mysql]-->初始化數據成功")
}

// var UserData = []User{
// 	{
// 		UserName: "jessie",
// 		Roles: []Role{
// 			{
// 				RoleName: "admin",
// 				Dashboards: []Dashboard{
// 					{
// 						DashboardName: "Single Node Analysis - cpu",
// 					},
// 					{
// 						DashboardName: "Single Node Analysis - mem",
// 					},
// 				},
// 			},
// 		},
// 	},
// }

// var InstanceData = []Instance{
// 	{
// 		Auth:         false,
// 		IP:           "http://10.99.1.240:3000",
// 		InstanceName: "grafana-240",
// 		Type:         "grafana",
// 		User:         "admin",
// 		Pass:         "12345678",
// 		Dashboards: []Dashboard{
// 			{
// 				DashboardName: "Single Node Analysis - cpu",
// 				Alias:         "Demo1",
// 				Description:   "cpu",
// 				UID:           "mOTd6X67k",
// 			},
// 			{
// 				DashboardName: "Single Node Analysis - mem",
// 				Alias:         "Demo2",
// 				Description:   "mem",
// 				UID:           "RlgkWFg4k",
// 			},
// 		},
// 	},
// }

// var MenuData = Menu{
// 	MenuName: "Single Node",
// 	Sort:     0,
// 	Dashboards: []Dashboard{
// 		{DashboardName: "Single Node Analysis - cpu"},
// 		{DashboardName: "Single Node Analysis - mem"},
// 	},
// }
