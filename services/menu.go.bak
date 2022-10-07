package services

import (
	"fmt"
	"sort"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

//get all menus include dashboards
func GetAllMenu() ([]models.Menu, error) {
	var dbData []entities.Menu
	var returnData []models.Menu

	err := global.Mysql.Preload("Dashboards").Find(&dbData).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(dbData); i++ {
		var menu models.Menu
		menu.MenuID = dbData[i].MenuID
		menu.MenuName = dbData[i].MenuName
		menu.Sort = dbData[i].Sort
		for j := 0; j < len(dbData[i].Dashboards); j++ {
			var dashboard models.Dashboard
			dashboard.DashboardID = dbData[i].Dashboards[j].DashboardID
			dashboard.DashboardName = dbData[i].Dashboards[j].DashboardName
			dashboard.Alias = dbData[i].Dashboards[j].Alias
			dashboard.Description = dbData[i].Dashboards[j].Description
			dashboard.UID = dbData[i].Dashboards[j].UID
			dashboard.MenuID = dbData[i].Dashboards[j].MenuID
			dashboard.InstanceID = dbData[i].Dashboards[j].InstanceID
			global.Mysql.Table("instances").Where("instance_id = ?", dbData[i].Dashboards[j].InstanceID).Select("instance_name").Scan(&dashboard.InstanceName)
			global.Mysql.Table("menus").Where("menu_id = ?", dbData[i].Dashboards[j].MenuID).Select("menu_name").Scan(&dashboard.MenuName)
			menu.Dashoboards = append(menu.Dashoboards, dashboard)
		}
		returnData = append(returnData, menu)
	}

	return returnData, nil
}

// get all dashboard of menu by menuID
func GetDashboardByMenuID(menuID int) ([]entities.Dashboard, error) {
	menu := entities.Menu{}
	menu.MenuID = menuID
	err := global.Mysql.Preload("Dashboards").Where("menu_id = ?", menuID).First(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu.Dashboards, nil
}

// create menu (create menu name only)
func CreateMenu(menu entities.Menu) models.Response {
	res := models.Response{}
	err := global.Mysql.Omit("Dashboards").Create(&menu).Error
	if err != nil {
		res.Msg = err.Error()
		res.Success = false
		return res
	}
	res.Msg = "Add new menu successfully"
	res.Success = true
	return res
}

// Check menu name
func CheckMenuName(menuName string) models.Response {

	res := models.Response{}
	rowsAffected := global.Mysql.Where("menu_name = ?", menuName).First(&entities.Menu{}).RowsAffected
	if rowsAffected == 0 {
		res.Msg = "ok"
		res.Success = true
		return res
	}
	res.Msg = "menu name existed"
	res.Success = false
	return res
}

// Check menu name in another id
func CheckMenuNameInOthers(menuID int, menuName string) models.Response {
	res := models.Response{}
	rowsAffected := global.Mysql.Where("menu_name = ? AND menu_id != ?", menuName, menuID).First(&entities.Menu{}).RowsAffected
	if rowsAffected == 0 {
		res.Msg = "ok"
		res.Success = true
		return res
	}
	res.Msg = "menu name existed in other menu id"
	res.Success = false
	return res
}

// Check menu by id
func CheckMenubyID(id int) models.Response {
	res := models.Response{}

	err := global.Mysql.Where("menu_id = ?", id).First(&entities.Menu{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when check menu ID, error: %s", err.Error())
		res.Success = false
		return res
	}

	res.Msg = "ok"
	res.Success = true
	return res
}

// add dashboard to menu
func AddMenuDashboard(menuID int, dashboard entities.Dashboard) models.Response {

	res := models.Response{}

	// 確認dashboard是否已存在DB中，若沒有則新增並加入Admin可以看的清單中
	check := CheckDashboard(menuID, dashboard)
	if !check.Success {
		res.Msg = check.Msg
		res.Success = false
		return res
	}

	// 新增 Dashboard (of the menuID)
	dashboard.MenuID = menuID
	err := global.Mysql.Create(&dashboard).Error
	if err != nil {
		res.Msg = fmt.Sprintf("error when creating dashboard to db, err: %v", err)
		res.Success = false
		return res
	}

	// 新增 Dashboard to Admin
	var dashboardID int
	global.Mysql.Table("dashboards").Where("uid = ? AND menu_id = ?", dashboard.UID, menuID).Select("dashboard_id").Scan(&dashboardID)
	err = global.Mysql.Table("dashboards_roles").Create(map[string]interface{}{"dashboard_id": dashboardID, "role_id": 1}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("error when creating dashboards_roles, err: %v", err)
		res.Success = false
		return res
	}

	res.Msg = "add dashboard successfully"
	res.Success = true
	return res
}

// update menu (update menu name only)
func UpdateMenu(menu entities.Menu) models.Response {
	res := models.Response{}
	err := global.Mysql.Omit("Dashboards").Where("menu_id = ?", menu.MenuID).Select("*").Updates(&menu).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "update menu successfully"
	res.Success = true
	return res
}

// delete menu (delete dashboard if the dashboard is not in other menu)
func DeleteMenuByID(menuID int) models.Response {
	menu := entities.Menu{}

	res := models.Response{}
	err := global.Mysql.Where("menu_id = ?", menuID).Delete(&menu).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}

	res.Msg = "Delete Success"
	res.Success = true
	return res
}

// delete dashboard from menu
func DeleteMenuDashboard(menuID, dashboardID int) models.Response {

	res := models.Response{}
	err := global.Mysql.Where("dashboard_id = ? AND menu_id = ?", dashboardID, menuID).Delete(&entities.Dashboard{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}

	res.Msg = "Delete successfully"
	res.Success = true
	return res
}

// get Base Menu
func GetDashboardMenuByID(userID int) ([]models.DashboardMenu, error) {
	var returnData []models.DashboardMenu

	// 用 userID 查詢有幾個 roleID
	var roles []int
	err := global.Mysql.Table("users_roles").Where("user_id = ?", userID).Select("role_id").Scan(&roles).Error
	if err != nil {
		return nil, err
	}

	// 用 roleID loop 查詢 user 可以看哪些 dashboard
	var dashboards, dashboardID_list []int
	for _, roleID := range roles {
		err := global.Mysql.Table("dashboards_roles").Where("role_id = ?", roleID).Select("dashboard_id").Scan(&dashboards).Error
		if err != nil {
			return nil, err
		}		
		dashboardID_list = append(dashboardID_list, dashboards...)
	}

	// 去除重複的 dashboardID
	dashboardID_list = RemoveRep(dashboardID_list)

	// 查詢 dashboardID 隸屬哪些 menu
	var menus, menuID_list []int
	for _, dashboardID := range dashboardID_list {
		err := global.Mysql.Table("dashboards").Where("dashboard_id = ?", dashboardID).Select("menu_id").Scan(&menus).Error
		if err != nil {
			return nil, err
		}
		menuID_list = append(menuID_list, menus...)
	}

	// 去除重複的 menuID
	menuID_list = RemoveRep(menuID_list)

	// 建立外層 menus 
	var parent models.DashboardMenu
	for _, menuID := range menuID_list {
		var menu entities.Menu
		err := global.Mysql.Table("menus").Where("menu_id = ?", menuID).Find(&menu).Error
		if err != nil {
			return nil, err
		}
		// menu detail
		parent.RouterPath = "main/active-module" // default 先寫死
		parent.ID = menu.MenuID
		parent.Title = menu.MenuName
		parent.Sort = menu.Sort
		returnData = append(returnData, parent)
	}
	
	// 用 dashboardID 查詢 dashboard detail
	var children models.DashboardMenu
	var dashboard_list []models.DashboardMenu
	for i, dashboardID := range dashboardID_list {
		var dashboard entities.Dashboard
		err := global.Mysql.Table("dashboards").Where("dashboard_id = ?", dashboardID).Find(&dashboard).Error
		if err != nil {
			return nil, err
		}
		// 用 InstanceID 找 type
		instanceID := dashboard.InstanceID
		var instance_type string
		err1 := global.Mysql.Table("instances").Where("instance_id = ?", instanceID).Select("type").Scan(&instance_type).Error
		if err1 != nil {
			return nil, err1
		}
		children.Type = instance_type
	
		children.RouterPath = "main/active-module" // default 先寫死
		children.ID = dashboard.DashboardID
		children.UID = dashboard.UID
		children.Title = dashboard.Alias // not DashboardName
		children.ParentID = dashboard.MenuID

		// dashboard detail 塞到 dashboard_list
		dashboard_list = append(dashboard_list, children)

		// 把 dashboard 塞到對應的 menu
		for j, menu := range returnData {
			if menu.ID ==  dashboard_list[i].ParentID {
				returnData[j].Children = append(returnData[j].Children, dashboard_list[i])
			}
		}	
	}
	return returnData, err
}

// 將 int list 去除重複的元素和 sort
func RemoveRep (numlist []int) ([]int) {
	result := []int{}
    for i := range numlist {
        flag := true
        for j := range result{
            if numlist[i] == result[j] {
                flag = false
                break
            }
        }
        if flag {
            result = append(result, numlist[i])
        }
    }
	sort.Ints(result)
	return result
}