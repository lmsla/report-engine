package services

import (
	"fmt"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

// 新增dashboards by role id
func AddRoleDashboard(roleID, dashboardID int) models.Response {
	res := models.Response{}

	err := global.Mysql.Table("dashboards_roles").Create(map[string]interface{}{"dashboard_id": dashboardID, "role_id": roleID}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}

	res.Msg = fmt.Sprintf("add dashboards to role %v success", roleID)
	res.Success = true
	return res
}

// 刪除dashboard by role id
func DeleteRoleDashboard(roleID, dashboardID int) models.Response {
	res := models.Response{}

	if roleID == 1 {
		delete := DeleteDashboard(dashboardID)
		if !delete.Success {
			return delete
		}
	} else {
		err := global.Mysql.Exec("DELETE FROM dashboards_roles WHERE dashboard_id = ? AND role_id = ?", dashboardID, roleID).Error
		if err != nil {
			res.Msg = fmt.Sprintf("Error: %v", err)
			res.Success = false
			return res
		}
	}

	res.Msg = fmt.Sprintf("Delete dashboard from role %v", roleID)
	res.Success = true
	return res
}

// Update dashboard
func UpdateRoleDashboard(roleID, oldID, newID int) models.Response {
	res := models.Response{}

	add := AddRoleDashboard(roleID, newID)
	if !add.Success {
		return add
	}
	delete := DeleteRoleDashboard(roleID, oldID)
	if !delete.Success {
		return delete
	}

	res.Msg = "Update Successfully"
	res.Success = true
	return res
}
