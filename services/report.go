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

	var body = []models.Report{}
	err := global.Mysql.Debug().Preload("Elements").Find(&body).Error

	if err != nil {
		res.Msg = err.Error()
		return res
	}
	res.Body = body
	res.Success = true
	res.Msg = "Get All Report Success"
	return res
}



// 查單一 report by ReportID
func GetReportByReportID(reportID int) (entities.Report,error) {

	var instance entities.Report
	instance.ID = reportID
	err := global.Mysql.Debug().Where("id = ?",reportID).Preload("Elements").Preload("Elements.Instance").Find(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil

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
	//先刪除 element 中相應的圖表
	err := global.Mysql.Where("report_id = ?", report.ID).Delete(&entities.Element{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting related elements, err: %s", err)
		return res
	}

	err = global.Mysql.Select("*").Where("id = ?", report.ID).Updates(&report).Error
	if err != nil {
		res.Msg = "Update Fail"
		return res
	}

	res.Success = true
	res.Msg = "Update Success"
	global.Mysql.Where("id = ?", report.ID).First(&res.Body)

	return res

}



func DeleteReport(id int) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = nil

	result := global.Mysql.Where("id = ?", id).First(&entities.Report{})
	if result.RowsAffected == 0 {
		res.Msg = "Report ID does not exist"
		return res
	}

	//先刪除 element 中相應的圖表
	err := global.Mysql.Where("report_id = ?", id).Delete(&entities.Element{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting related elements, err: %s", err)
		return res
	}
	//刪除 ReportsSchedule 中對應的 Report
	err = global.Mysql.Where("report_id = ?" ,id).Delete(&entities.ReportsSchedules{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting data in ReportsSchedule, err: %s", err)
		return res
	}

	err = global.Mysql.Where("id = ?", id).Delete(&entities.Report{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting report, err: %s", err)
		return res
	}

	res.Success = true
	res.Msg = "Delete Success"

	return res

}


// // 查 Reports by ScheduleID
// func GetReportByScheduleID(scheduleID int) (models.Response) {

// 	// many2many檢索
// 	res := models.Response{}
// 	res.Success = false
// 	var schedule entities.Schedule
// 	// var body = []entities.Schedule{}

// 	err := global.Mysql.Debug().Where("id = ?",scheduleID).Preload("Reports").Find(&schedule).Error

// 	// err := global.Mysql.Debug().Where("schedule_id = ?",scheduleID).Preload("Report").Find(&body).Error
// 	if err != nil {
// 		res.Msg = err.Error()
// 		return res
// 	}

// 	res.Body = schedule.Reports
// 	res.Success = true
// 	res.Msg = "Get Selected Report Success"
// 	return res


// }


func GetReportByScheduleID(scheduleID int) ([]entities.Report,error) {

	// // many2many檢索
	// res := models.Response{}
	// res.Success = false
	// var schedule entities.Schedule
	// var body = []entities.Schedule{}
	schedule := entities.Schedule{}
	// err := global.Mysql.Debug().Where("id = ?",scheduleID).Preload("Reports").Find(&schedule).Error

	err := global.Mysql.Debug().Where("id = ?",scheduleID).Preload("Reports").Find(&schedule).Error
	if err != nil {
		return nil,err
	}
	return schedule.Reports,nil

}