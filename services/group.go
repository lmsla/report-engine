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

// 確認 ID 是否存在
func GetGroupByID(id int) ([]entities.Group, error) {
	columns := []entities.Group{}
	err := global.Mysql.Where("group_id = ?", id).Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}


// 確認 Name 是否重複
func GetGroupByName(name string) ([]entities.Group, error) {
	columns := []entities.Group{}

	err := global.Mysql.Where("name = ?", name).Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}
// 確認 Member ID 是否存在
func GetMemberByID(id int) ([]entities.GroupMember, error) {
	columns := []entities.GroupMember{}
	err := global.Mysql.Where("member_id = ?", id).Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}


// 確認 MemberName 是否重複
func GetMemberByName(name string) ([]entities.GroupMember, error) {
	columns := []entities.GroupMember{}

	err := global.Mysql.Where("member_name = ?", name).Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}


// 更新group
func UpdateGroup(r entities.Group) models.Response {
	res := models.Response{}

	// var member entities.GroupMember
	temp := entities.Group{}
	DBresponse := global.Mysql.Where("group_id = ?", r.GroupID).First(&temp)

	if DBresponse.RowsAffected == 0 {
		res.Success = false
		res.Msg = "Group id does not exist"
		return res
	}

	r.CreatedAt = temp.CreatedAt

	err := global.Mysql.Select("*").Where("group_id = ?", r.GroupID).Updates(&r).Error
	if err != nil {
		res.Success = false
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = fmt.Sprintf("Group ID %v Updated Success",  r.GroupID)
	return res
}


// 更新 member
func UpdateMember1(r entities.GroupMember) models.Response {
	res := models.Response{}

	// 1. Update users table
	var member entities.GroupMember
	member.MemberID = r.MemberID
	member.MemberName = r.MemberName
	err1 := global.Mysql.Where("member_id = ?", r.MemberID).Updates(&member).Error 
	if err1 != nil {
		res.Msg = fmt.Sprintf("Error: %v", err1)
		res.Success = false
		return res
	}

	res.Msg = "Update Success"
	res.Success = true
	return res
}


// 更新member
func UpdateMember(r entities.GroupMember) models.Response {
	res := models.Response{}
	// var member entities.GroupMember
	temp := entities.GroupMember{}
	DBresponse := global.Mysql.Where("member_id = ?", r.MemberID).First(&temp)

	if DBresponse.RowsAffected == 0 {
		res.Success = false
		res.Msg = "Member id does not exist"
		return res
	}

	r.CreatedAt = temp.CreatedAt

	err := global.Mysql.Select("*").Where("member_id = ?", r.MemberID).Updates(&r).Error
	if err != nil {
		res.Success = false
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = fmt.Sprintf("Member ID %v Updated Success",  r.MemberID)
	return res
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


// 刪除 Group by id
func DeleteGroup(groupID int) models.Response {

	res := models.Response{}
	DBresponse := global.Mysql.Where("group_id = ?", groupID).Delete(&entities.Group{})

	if DBresponse.RowsAffected == 0 {
		res.Msg = fmt.Sprintf("groupID %v does not exist", groupID)
		res.Success = false
		return res
	}
	if DBresponse.Error != nil {
		res.Msg = fmt.Sprintf("Error: %v", DBresponse.Error)
		res.Success = false
		return res
	}
	res.Msg = fmt.Sprintf("groupID %v Deleted", groupID)
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



