package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

func GetAllInstances() models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []models.Instance{}

	err := global.Mysql.Find(&res.Body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = "Get All Instance Success"
	return res
}

// 新增instance
func CreateInstance(instance entities.Instance) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []models.Instance{}

	result := global.Mysql.Where("name = ?", instance.Name).First(&entities.Instance{})
	if result.RowsAffected > 0 {
		res.Msg = "Instance Name already existed"
		return res
	}

	err := global.Mysql.Create(&instance).Error
	if err != nil {
		res.Msg = "Create Fail"
		return res
	}

	res.Success = true
	res.Msg = "Create Success"
	global.Mysql.Where("name = ?", instance.Name).First(&res.Body)

	return res
}


// 查單一Instance
func GetInstanceByID(instanceID int) (models.Instance, error) {

	var instance models.Instance
	instance.ID = instanceID
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}

func UpdateInstance(instance models.Instance) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []models.Instance{}

	result := global.Mysql.Where("id != ? AND name = ?", instance.ID, instance.Name).First(&entities.Instance{})
	if result.RowsAffected > 0 {
		res.Msg = "Instance Name already existed"
		return res
	}

	err := global.Mysql.Select("*").Where("id = ?", instance.ID).Updates(&instance).Error
	if err != nil {
		res.Msg = "Update Fail"
		return res
	}

	res.Success = true
	res.Msg = "Update Success"
	global.Mysql.Where("id = ?", instance.ID).First(&res.Body)

	return res

}


func DeleteInstance(id int) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = nil

	result := global.Mysql.Where("id = ?", id).First(&entities.Instance{})
	if result.RowsAffected == 0 {
		res.Msg = "Instance ID does not exist"
		return res
	}

	//先刪除 element 中相應的圖表
	err := global.Mysql.Where("instance_id = ?", id).Delete(&entities.Element{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting related elements, err: %s", err)
		return res
	}

	err = global.Mysql.Where("id = ?", id).Delete(&entities.Instance{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting instance, err: %s", err)
		return res
	}

	res.Success = true
	res.Msg = "Delete Success"

	return res

}


