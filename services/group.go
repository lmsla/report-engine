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
