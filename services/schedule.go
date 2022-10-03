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


// 更新Schedule
func UpdateSchedule1(scheduleID int, instance entities.Schedule) models.Response {

	res := models.Response{}

	temp := entities.Schedule{}
	// temp := entities.Status{}
	DBresponse := global.Mysql.Where("schedule_id = ?", scheduleID).First(&temp)

	if DBresponse.RowsAffected == 0 {
		res.Success = false
		res.Msg = "schedule id does not exist"
		return res
	}

	instance.CreatedAt = temp.CreatedAt
	// 註：.Select("*") 會導致 struct 中沒更新的欄位(沒填值)直接消失，此處只需更新 Status，故拿掉 .Select("*") ，API 送進去的 struct 可以只更新 status 
	// err := global.Mysql.Select("*").Where("schedule_id = ?", scheduleID).Updates(&instance).Error
	err := global.Mysql.Where("schedule_id = ?", scheduleID).Updates(&instance).Error
	if err != nil {
		res.Success = false
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = fmt.Sprintf("schedule ID %v Updated Success", scheduleID)
	return res
}


// 更新Schedule
func UpdateSchedule(schedule entities.Schedule) models.Response {

	res := models.Response{}

	temp := entities.Schedule{}
	// temp := entities.Status{}
	DBresponse := global.Mysql.Where("schedule_id = ?", schedule.ScheduleID).First(&temp)

	if DBresponse.RowsAffected == 0 {
		res.Success = false
		res.Msg = "schedule id does not exist"
		return res
	}

	schedule.CreatedAt = temp.CreatedAt
	// 註：.Select("*") 會導致 struct 中沒更新的欄位(沒填值)直接消失，此處只需更新 Status，故拿掉 .Select("*") ，API 送進去的 struct 可以只更新 status 
	// err := global.Mysql.Select("*").Where("schedule_id = ?", scheduleID).Updates(&instance).Error
	err := global.Mysql.Where("schedule_id = ?", schedule.ScheduleID).Updates(&schedule).Error
	if err != nil {
		res.Success = false
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = fmt.Sprintf("schedule ID %v Updated Success", schedule.ScheduleID)
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

	ScreenshotDocker1(inventory1)
	
	CreateHtml1(scheduleID,inventory1)

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
		Sendmail(scheduleID,inventory.Gropup)

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
		if inventory.Status == "run" {
			global.Crontab.Start()
		}
		// global.Crontab.Start()

	}

}

// 刪除 Schedule by id
func DeleteSchedule(scheduleID int) models.Response {

	res := models.Response{}
	DBresponse := global.Mysql.Where("schedule_id = ?", scheduleID).Delete(&entities.Schedule{})

	if DBresponse.RowsAffected == 0 {
		res.Msg = fmt.Sprintf("scheduleID %v does not exist", scheduleID)
		res.Success = false
		return res
	}
	if DBresponse.Error != nil {
		res.Msg = fmt.Sprintf("Error: %v", DBresponse.Error)
		res.Success = false
		return res
	}
	res.Msg = fmt.Sprintf("scheduleID %v Deleted", scheduleID)
	res.Success = true
	return res
}