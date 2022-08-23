package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

// Add Dashboard
func CreateDashboard(dashboard entities.Dashboard) models.Response {
	res := models.Response{}
	err := global.Mysql.Create(&dashboard).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "create dashboard successfully"
	res.Success = true
	return res
}

// Check Dashboard
func CheckDashboard(menuID int, dashboard entities.Dashboard) models.Response {
	res := models.Response{}
	dbResponse := global.Mysql.Where("uid = ? AND menu_id = ?", dashboard.UID, menuID).First(&entities.Dashboard{})
	if dbResponse.RowsAffected != 0 {
		res.Msg = "dashboard already existed"
		res.Success = false
		return res
	}
	res.Msg = "ok"
	res.Success = true
	return res
}

// Check Dashboard ID
func CheckDashboardID(dashboardID int) models.Response {
	res := models.Response{}
	dbResponse := global.Mysql.Where("dashboard_id", dashboardID).First(&entities.Dashboard{})
	if dbResponse.RowsAffected == 0 {
		res.Msg = "dashboard does not exist"
		res.Success = false
		return res
	}
	res.Msg = "ok"
	res.Success = true
	return res
}

// // Check Dashbaord in menus_dashboards
// func CheckMenuDashboard(dashboardID int) models.Response {
// 	res := models.Response{}
// 	var menuID int
// 	dbResponse := global.Mysql.Debug().Table("menus_dashboards").Where("dashboard_id = ?", dashboardID).Select("menu_id").Scan(&menuID)
// 	if dbResponse.RowsAffected == 0 {
// 		res.Msg = "No more related menu to this dashboard"
// 		res.Success = false
// 		return res
// 	}
// 	res.Msg = "the dashboard still under other menu"
// 	res.Success = true
// 	return res
// }

// Get Dashboards by Instance
func GetDBDashboardsbyInstanceID(instanceID int) ([]entities.Dashboard, error) {
	dashboards := []entities.Dashboard{}
	err := global.Mysql.Where("instance_id = ?", instanceID).Find(&dashboards).Error
	if err != nil {
		return nil, err
	}
	return dashboards, nil
}

// Delete Dashboard
func DeleteDashboard(dashboardID int) models.Response {
	res := models.Response{}
	err := global.Mysql.Where("dashboard_id = ?", dashboardID).Delete(&entities.Dashboard{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("fail to delete dashboard from DB, error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Delete Successfully"
	res.Success = true
	return res
}

// Update Dashboard
func UpdateDashboard(model models.Dashboard) models.Response {
	res := models.Response{}
	fmt.Println("description: ", model.Description, "alias: ", model.Alias)
	err := global.Mysql.Debug().Table("dashboards").Where("dashboard_id = ?", model.DashboardID).Updates(map[string]interface{}{"alias": model.Alias, "description": model.Description}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("error when update info of dashboard, error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Update dashboard alias and description successfully"
	res.Success = true
	return res
}

// Check Dashboard Alias in the same menu
func CheckAliasInOthers(dashboardID, menuID int, alias string) models.Response {
	res := models.Response{}
	dbResponse := global.Mysql.Where("dashboard_id != ? AND menu_id = ? AND alias = ?", dashboardID, menuID, alias).First(&entities.Dashboard{})
	if dbResponse.RowsAffected != 0 {
		res.Msg = fmt.Sprintf("alias name already existed under menu id %v", menuID)
		res.Success = false
		return res
	}
	res.Msg = "ok"
	res.Success = true
	return res
}
