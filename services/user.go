package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

func GetUserAll() ([]entities.User, error) {
	columns := []entities.User{}
	err := global.Mysql.Preload("Roles").Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}

// 確認 ID 是否存在
func GetUserByID(id int) ([]entities.User, error) {
	columns := []entities.User{}
	err := global.Mysql.Where("user_id = ?", id).Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}

// 確認 Name 是否重複
func GetUserByName(name string) ([]entities.User, error) {
	columns := []entities.User{}

	err := global.Mysql.Where("user_name = ?", name).Find(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}

func CreateUser(r entities.User) models.Response {
	res := models.Response{}

	// 1. Create users table
	var user entities.User
	user.UserID = r.UserID
	user.UserName = r.UserName

	err := global.Mysql.Omit("Roles").Create(&user).Error // 先略過Role的欄位
	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}

	// 2. Create users_roles table
	users, _ := GetUserByName(r.UserName)
	users_roles := entities.UsersRole{}

	for _, role := range r.Roles {
		users_roles.UserID = users[0].UserID
		users_roles.RoleID = role.RoleID

		err := global.Mysql.Create(&users_roles).Error
		if err != nil {
			res.Msg = fmt.Sprintf("Error: %v", err)
			res.Success = false
			return res
		}
	}
	res.Msg = "Create Success"
	res.Success = true
	return res
}

func UpdateUser(r entities.User) models.Response {
	res := models.Response{}

	// 1. Update users table
	var user entities.User
	user.UserID = r.UserID
	user.UserName = r.UserName
	err1 := global.Mysql.Omit("Roles").Where("user_id = ?", r.UserID).Updates(&user).Error // 先略過Role的欄位
	if err1 != nil {
		res.Msg = fmt.Sprintf("Error: %v", err1)
		res.Success = false
		return res
	}

	// 2. 先 Delete 原本的 users_roles table
	users_roles := entities.UsersRole{}

	err2 := global.Mysql.Where("user_id = ?", r.UserID).Delete(&users_roles).Error
	if err2 != nil {
		res.Msg = fmt.Sprintf("Error: %v", err2)
		res.Success = false
		return res
	}

	// 3. Create 新的 users_roles table
	for _, role := range r.Roles {
		users_roles.UserID = r.UserID
		users_roles.RoleID = role.RoleID

		err := global.Mysql.Create(&users_roles).Error
		if err != nil {
			res.Msg = fmt.Sprintf("Error: %v", err)
			res.Success = false
			return res
		}
	}
	res.Msg = "Update Success"
	res.Success = true
	return res
}

func DeleteUserByID(id int) models.Response {
	user := []entities.User{}
	res := models.Response{}

	err := global.Mysql.Where("user_id = ?", id).Delete(&user).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Delete Success"
	res.Success = true
	return res
}
