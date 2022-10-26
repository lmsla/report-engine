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
	err := global.Mysql.Preload("Dashboard").Preload("Instance").Find(&instancies).Error
	// err := global.Mysql.Find(&instancies).Error
	if err != nil {
		return nil, err
	}
	return instancies, nil
}


// 查單一Report
func GetReportByReportID(reportID int) (entities.Report, error) {

	var instance entities.Report
	instance.ReportID = reportID
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}

// 藉由report name查report
func GetReportByReportName([]entities.Report) ([]entities.Report, error) {

	var instance []entities.Report
	// instance.Name = reportName
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}


// 藉由report name查report
func GetReportByReportName1(reportName string) (entities.Report, error) {

	var instance entities.Report
	instance.Name = reportName
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}


// 查report中有哪些instance
func GetInstanceInReportbyReportID(reportID int)([]entities.RInstance,error) {
	instance := entities.Report{}
	instance.ReportID = reportID
	// err := global.Mysql.Find(&instance).Error
	err := global.Mysql.Preload("Instance").Where("report_id = ?", reportID).Find(&instance).Error
	if err != nil {
		return nil, err
	}
	return instance.Instance, nil

}

// 取出 report中的dashboard
func GetDashboardInReport(reportID int) ([]entities.RDashboard, error){
	// var IPType []map[string]int
	instance := entities.Report{}
	instance.ReportID = reportID
	// err := global.Mysql.Find(&instance).Error
	err := global.Mysql.Preload("Dashboard").Where("report_id = ?", reportID).Find(&instance).Error
	if err != nil {
		return nil, err
	}
	return instance.Dashboard, nil

}


// 刪除Report
func DeleteReport(reportID int) models.Response {

	res := models.Response{}
	DBresponse := global.Mysql.Where("report_id = ?", reportID).Delete(&entities.Report{})

	if DBresponse.RowsAffected == 0 {
		res.Msg = fmt.Sprintf("Report ID %v does not exist", reportID)
		res.Success = false
		return res
	}
	if DBresponse.Error != nil {
		res.Msg = fmt.Sprintf("Error: %v", DBresponse.Error)
		res.Success = false
		return res
	}
	res.Msg = fmt.Sprintf("Report ID %v Deleted", reportID)
	res.Success = true
	return res
}


// 取出dashboard by instance ID 
func GetDashboardOfInstanceByInstanceID(instanceID int) ([]entities.RDashboard, error){
	instance := entities.RInstance{}
	instance.InstanceID = instanceID
	err := global.Mysql.Preload("Dashboard").Where("instance_id = ?", instanceID).Find(&instance).Error
	if err != nil {
		return nil, err
	}
	return instance.Dashboard, nil

}




// // 取出 report中的dashboard (未定)
// func GetDashboardInReport1(reportID int) ([]map[string]interface{}, error){
// 	// var IPType []map[string]int
// 	var instance entities.RDashboard
// 	instance.ReportID = reportID
// 	err := global.Mysql.First(&instance).Error
	
// 	var DashboardList []map[string]interface{}
// 	if err != nil {
// 		// return instance, err
// 		var rawdata map[string]interface{}
// 		rawdata = make(map[string]interface{})
// 		rawdata["UID"] = instance.UID
// 		rawdata["instanceID"] = instance.InstanceID
// 		rawdata["reportID"] = instance.ReportID
// 		DashboardList = append(DashboardList, rawdata)

// 	}
// 	return DashboardList, nil

// }