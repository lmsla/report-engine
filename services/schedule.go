package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
	// "report-backend-golang/screenshot"
	// "report-backend-golang/schedule"
	// "github.com/robfig/cron/v3"
	"report-backend-golang/log"
)

// 新增 schedule
func CreateSchedule(schedule entities.Schedule) models.Response {

	res := models.Response{}
	err := global.Mysql.Create(&schedule).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	} else {
		fmt.Println(schedule.ScheduleID)
		// 排程新增後自動 add func to cron
		ExecuteSheduleSendMail(schedule.ScheduleID)
		ExecuteShedulePDF(schedule.ScheduleID)
	}
	res.Msg = "Create Success"
	res.Success = true
	return res

}

// 查詢All schedule
func GetAllSchedule() ([]entities.Schedule, error) {

	var instancies []entities.Schedule
	err := global.Mysql.Find(&instancies).Error
	if err != nil {
		return nil, err
	}
	return instancies, nil
}

// 查單一Schedule
func GetScheduleBysSheduleID(scheduleID int) (entities.Schedule, error) {

	var instance entities.Schedule
	instance.ScheduleID = scheduleID
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}



func FuncAddToCron(scheduleID int) {
	inventory,err := GetScheduleBysSheduleID(scheduleID)
	if err != nil {
		fmt.Println(err)
	}
	inventory1,err := GetReportByReportName(inventory.Report)

	Screenshot(inventory1.ReportID)
	
	CreateHtml(scheduleID,inventory1.ReportID)

	// Sendmail(scheduleID,inventory.Recipient)

}




// 執行 PDF Schedule by ScheduleID
func ExecuteShedulePDF(scheduleID int) {
	inventory,err := GetScheduleBysSheduleID(scheduleID)
	if err != nil {
		fmt.Println(err)
	}
	// inventory1,err := services.GetReportByReportName(inventory.Report)

	_,err = global.Crontab.AddFunc(inventory.GenerateReport,func(){
		FuncAddToCron(scheduleID)

	}) 
	//fmt.Print(global.EnvConfig.CRONTAB.Period,global.EnvConfig.INFLUX.URL)
	if err != nil {
		fmt.Println("crontab PDF 初始化失敗")
		log.Logrecord("排程 ","PDF排程 初始化失敗")
		fmt.Println(err.Error())
		log.Logrecord("ERROR ",err.Error())
	} else {
		fmt.Println("crontab PDF 初始化成功")
		log.Logrecord("排程 ","PDF排程 初始化成功")
		// c.Start()
		global.Crontab.Start()

	}

}


// 執行 SendMail Schedule by ScheduleID
func ExecuteSheduleSendMail(scheduleID int) {
	inventory,err := GetScheduleBysSheduleID(scheduleID)
	if err != nil {
		fmt.Println(err)
	}
	// inventory1,err := services.GetReportByReportName(inventory.Report)

	_,err = global.Crontab.AddFunc(inventory.SendReport,func(){
		Sendmail(scheduleID,inventory.Recipient)

	}) 
	//fmt.Print(global.EnvConfig.CRONTAB.Period,global.EnvConfig.INFLUX.URL)
	if err != nil {
		fmt.Println("crontab xdr 初始化失敗")
		log.Logrecord("排程 ","xdr排程 初始化失敗")
		fmt.Println(err.Error())
		log.Logrecord("ERROR ",err.Error())
	} else {
		fmt.Println("crontab xdr 初始化成功")
		log.Logrecord("排程 ","xdr排程 初始化成功")
		// c.Start()
		global.Crontab.Start()

	}

}