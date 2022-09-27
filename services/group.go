package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

// 新增 group
func CreateGroup(group entities.Group) models.Response {

	res := models.Response{}
	err := global.Mysql.Create(&group).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Create Success"
	res.Success = true
	return res
}

// 查詢All group
func GetAllGroup() ([]entities.Group, error) {

	var instancies []entities.Group
	err := global.Mysql.Preload("GroupMember").Find(&instancies).Error
	if err != nil {
		return nil, err
	}
	return instancies, nil
}

// 新增 member
func CreateMember(group entities.GroupMember) models.Response {

	res := models.Response{}
	err := global.Mysql.Create(&group).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Create Success"
	res.Success = true
	return res
}

// add dashboard to menu
func AddMemberToGroup(GroupID int, member entities.GroupMember) models.Response {

	res := models.Response{}

	// // 確認dashboard是否已存在DB中，若沒有則新增並加入Admin可以看的清單中
	// check := CheckDashboard(GroupID, member)
	// if !check.Success {
	// 	res.Msg = check.Msg
	// 	res.Success = false
	// 	return res
	// }

	// 新增 Dashboard (of the menuID)
	// member.GroupID = groupID
	err := global.Mysql.Create(&member).Error
	if err != nil {
		res.Msg = fmt.Sprintf("error when creating dashboard to db, err: %v", err)
		res.Success = false
		return res
	}

	// // 新增 Dashboard to Admin
	// var dashboardID int
	// global.Mysql.Table("dashboards").Where("uid = ? AND menu_id = ?", dashboard.UID, menuID).Select("dashboard_id").Scan(&dashboardID)
	// err = global.Mysql.Table("dashboards_roles").Create(map[string]interface{}{"dashboard_id": dashboardID, "role_id": 1}).Error
	// if err != nil {
	// 	res.Msg = fmt.Sprintf("error when creating dashboards_roles, err: %v", err)
	// 	res.Success = false
	// 	return res
	// }

	res.Msg = "add member successfully"
	res.Success = true
	return res
}






// 刪除 member by id
func DeleteMember(memberID int) models.Response {

	res := models.Response{}
	DBresponse := global.Mysql.Where("member_id = ?", memberID).Delete(&entities.GroupMember{})

	if DBresponse.RowsAffected == 0 {
		res.Msg = fmt.Sprintf("memberID %v does not exist", memberID)
		res.Success = false
		return res
	}
	if DBresponse.Error != nil {
		res.Msg = fmt.Sprintf("Error: %v", DBresponse.Error)
		res.Success = false
		return res
	}
	res.Msg = fmt.Sprintf("memberID %v Deleted", memberID)
	res.Success = true
	return res
}

// 刪除 member by name
func DeleteMemberbyName(memberName string) models.Response {

	res := models.Response{}
	DBresponse := global.Mysql.Where("member_name = ?", memberName).Delete(&entities.GroupMember{})

	if DBresponse.RowsAffected == 0 {
		res.Msg = fmt.Sprintf("memberName %v does not exist", memberName)
		res.Success = false
		return res
	}
	if DBresponse.Error != nil {
		res.Msg = fmt.Sprintf("Error: %v", DBresponse.Error)
		res.Success = false
		return res
	}
	res.Msg = fmt.Sprintf("memberName %v Deleted", memberName)
	res.Success = true
	return res
}


// // 查 group mamber by groupID
// func GetMemberByGroupID(GroupID int) ([]entities.GroupMember, error) {

// 	var instancies []entities.GroupMember
// 	instancies.GroupID = GroupID
// 	err := global.Mysql.Preload("GroupMember").Find(&instancies).Error
// 	// err := global.Mysql.First(&instance).Error

// 	if err != nil {
// 		return instancies, err
// 	}
// 	return instancies, nil
// }


// 查 group member by groupID
func GetMemberByGroupID(groupID int) ([]entities.GroupMember, error){
	// var IPType []map[string]int
	instance := entities.Group{}
	instance.GroupID = groupID
	// err := global.Mysql.Find(&instance).Error
	err := global.Mysql.Preload("GroupMember").Where("group_id = ?", groupID).Find(&instance).Error
	if err != nil {
		return nil, err
	}
	return instance.GroupMember, nil

}

// 查 group member by groupName
func GetMemberByGroupName(groupName string) ([]entities.GroupMember, error){
	// var IPType []map[string]int
	instance := entities.Group{}
	instance.Name = groupName
	// err := global.Mysql.Find(&instance).Error
	err := global.Mysql.Preload("GroupMember").Where("name = ?", groupName).Find(&instance).Error
	if err != nil {
		return nil, err
	}
	return instance.GroupMember, nil

}

// Check Group by id
func CheckGroupID(id int) models.Response {
	res := models.Response{}

	err := global.Mysql.Where("group_id = ?", id).First(&entities.Group{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when check group ID, error: %s", err.Error())
		res.Success = false
		return res
	}

	res.Msg = "ok"
	res.Success = true
	return res
}



