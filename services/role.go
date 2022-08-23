package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

// 取得所有Role包含Dashboard
func GetRoleAll() ([]models.Role, error) {
	dbData := []entities.Role{}
	returnData := []models.Role{}

	err := global.Mysql.Preload("Dashboards").Find(&dbData).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(dbData); i++ {
		var role models.Role
		role.RoleID = dbData[i].RoleID
		role.RoleName = dbData[i].RoleName
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
			role.Dashboards = append(role.Dashboards, dashboard)
		}
		returnData = append(returnData, role)
	}

	return returnData, nil
}

// 確認 ID 是否存在
func CheckRoleById(id int) models.Response {
	res := models.Response{}

	err := global.Mysql.Where("role_id = ?", id).First(&entities.Role{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when check Role ID, error: %s", err.Error())
		res.Success = false
		return res
	}

	res.Msg = "ok"
	res.Success = true
	return res
}

// 確認 Name 是否重複
func CheckRoleByNameInOther(roleID int, name string) models.Response {
	res := models.Response{}

	dbResponse := global.Mysql.Where("role_name = ? AND role_id != ?", name, roleID).First(&entities.Role{})
	if dbResponse.RowsAffected != 0 {
		res.Msg = fmt.Sprintln("Role Name already existed")
		res.Success = false
		return res
	}

	res.Msg = "ok"
	res.Success = true
	return res
}

// 確認 Name 是否重複
func CheckRoleByName(name string) models.Response {
	res := models.Response{}

	dbResponse := global.Mysql.Where("role_name = ?", name).First(&entities.Role{})
	if dbResponse.RowsAffected != 0 {
		res.Msg = fmt.Sprintln("Role Name already existed")
		res.Success = false
		return res
	}

	res.Msg = "ok"
	res.Success = true
	return res
}

// 新增role(不包含Dashboard)
func CreateRole(role entities.Role) models.Response {
	//records [][]string,
	res := models.Response{}
	err := global.Mysql.Create(&role).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Create Success"
	res.Success = true
	return res
}

// 更新role(不包含Dashboard)
func UpdateRole(r entities.Role) models.Response {
	res := models.Response{}
	err := global.Mysql.Debug().Where("role_id = ?", r.RoleID).Updates(&r).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Update Success"
	res.Success = true
	return res
}

// 刪除Role（同時刪除Relation Table - dashboards_roles)
func DeleteRoleById(id int) models.Response {
	Role := []entities.Role{}
	res := models.Response{}

	err := global.Mysql.Where("role_id = ?", id).Delete(&Role).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}

	res.Msg = "Delete Success"
	res.Success = true
	return res
}
