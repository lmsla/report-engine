package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

// 新增instance
func CreateInstance(instance entities.Instance) models.Response {

	res := models.Response{}
	// err := global.Mysql.Create(&instance).Error
	err := global.Mysql.Create(&instance).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Create Success"
	res.Success = true
	return res
}

// 查詢All Instance
func GetAllInstance() ([]entities.Instance, error) {

	var instancies []entities.Instance
	err := global.Mysql.Omit("Dashboards").Find(&instancies).Error
	if err != nil {
		return nil, err
	}
	return instancies, nil
}

// 查單一Instance
func GetInstanceByID(instanceID int) (entities.Instance, error) {

	var instance entities.Instance
	instance.InstanceID = instanceID
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}

// 更新Instance
func UpdateInstance(instanceID int, instance entities.Instance) models.Response {

	res := models.Response{}

	temp := entities.Instance{}
	DBresponse := global.Mysql.Where("instance_id = ?", instanceID).First(&temp)

	if DBresponse.RowsAffected == 0 {
		res.Success = false
		res.Msg = "Instance id does not exist"
		return res
	}

	instance.CreatedAt = temp.CreatedAt

	err := global.Mysql.Select("*").Where("instance_id = ?", instanceID).Updates(&instance).Error
	if err != nil {
		res.Success = false
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = fmt.Sprintf("Instance ID %v Updated Success", instanceID)
	return res
}

// 刪除Instance
func DeleteInstance(instanceID int) models.Response {

	res := models.Response{}
	DBresponse := global.Mysql.Where("instance_id = ?", instanceID).Delete(&entities.Instance{})

	if DBresponse.RowsAffected == 0 {
		res.Msg = fmt.Sprintf("Instacne ID %v does not exist", instanceID)
		res.Success = false
		return res
	}
	if DBresponse.Error != nil {
		res.Msg = fmt.Sprintf("Error: %v", DBresponse.Error)
		res.Success = false
		return res
	}
	res.Msg = fmt.Sprintf("Instacne ID %v Deleted", instanceID)
	res.Success = true
	return res
}
