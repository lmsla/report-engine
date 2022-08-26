package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

// 新增報表
func CreateReport(report entities.Report) models.Response {

	res := models.Response{}
	err := global.Mysql.Create(&report).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Msg = "Create Success"
	res.Success = true
	return res
}

// 查詢所有報表
func GetAllReport() ([]entities.Report, error) {

	var instancies []entities.Report
	err := global.Mysql.Find(&instancies).Error
	if err != nil {
		return nil, err
	}
	return instancies, nil
}


// 查單一Instance
func GetInstanceByReportID(reportID int) (entities.Report, error) {

	var instance entities.Report
	instance.ReportID = reportID
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}