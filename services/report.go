package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

func GetAllReports() models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []models.Report{}

	err := global.Mysql.Find(&res.Body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = "Get All Report Success"
	return res
}



func CreateReport(report models.Report) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []models.Report{}

	result := global.Mysql.Where("name = ?", report.Name).First(&models.Report{})
	if result.RowsAffected > 0 {
		res.Msg = "Report Name already existed"
		return res
	}

	err := global.Mysql.Create(&report).Error
	if err != nil {
		res.Msg = "Create Fail"
		return res
	}

	res.Success = true
	res.Msg = "Create Success"
	global.Mysql.Where("name = ?", report.Name).First(&res.Body)

	return res
}

func UpdateReport(report models.Report) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []models.Report{}

	result := global.Mysql.Where("id != ? AND name = ?", report.ID, report.Name).First(&models.Report{})
	if result.RowsAffected > 0 {
		res.Msg = "Report Name already existed"
		return res
	}

	err := global.Mysql.Select("*").Where("id = ?", report.ID).Omit("Elements").Updates(&report).Error
	if err != nil {
		res.Msg = "Update Fail"
		return res
	}

	res.Success = true
	res.Msg = "Update Success"
	global.Mysql.Where("id = ?", report.ID).First(&res.Body)

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