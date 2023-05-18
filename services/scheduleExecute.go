package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	// "report-backend-golang/models"
	// "report-backend-golang/screenshot"
	// "report-backend-golang/schedule"
	// "github.com/robfig/cron/v3"
	"report-backend-golang/log"
	// "time"
)

// 執行 PDF Schedule by ScheduleID
func ExecuteShedulePDF(scheduleID int) {

	inventory,err := GetScheduleBysSheduleID(scheduleID)
	if err != nil {
		fmt.Println(err)
	}
	history := entities.History{}
	history.ScheduleID = scheduleID
	history.To = inventory.To
	fmt.Println(history.ScheduleID)
	fmt.Println(history)

	// inventory1,err := services.GetReportByReportName(inventory.Report)

	// inventory.CronID
	EntryID,err := global.Crontab.AddFunc(inventory.CronTime,func(){
		fmt.Println("執行排程")
		err := global.Mysql.Create(&history).Error
		if err != nil {

		}
		FuncAddToCron(scheduleID)
	}) 
	fmt.Println("entryID: ")
	fmt.Println(EntryID,err)

	// 寫一筆記錄到 cron_lists 的 table 中
	// res := models.Response{}
	// res.Success = false
	cronlist := entities.CronList{ScheduleID: scheduleID,EntryID: int(EntryID)}
	result  := global.Mysql.Create(&cronlist).Error
	if result != nil {
		fmt.Println("Create Fail")
		// return res) 
	}

	// res.Success = true
	// res.Msg = "Create Success"

	//fmt.Print(global.EnvConfig.CRONTAB.Period,global.EnvConfig.INFLUX.URL)
	if err != nil {
		fmt.Println("crontab PDF 初始化失敗")
		// log.Logrecord("排程 ","PDF排程 初始化失敗")
		fmt.Println(err.Error())
		log.Logrecord("ERROR ","PDF排程 初始化失敗"+err.Error())
	} else {

		fmt.Println("crontab PDF 初始化成功")
		log.Logrecord("排程 ","PDF排程 初始化成功")
		// c.Start()
		global.Crontab.Start()

	}

}






func FuncAddToCron(scheduleID int) {
	inventory,err := GetScheduleBysSheduleID(scheduleID)
	if err != nil {
		log.Logrecord("ERROR ","Get Schedule by Schedule Id error" + err.Error())
		fmt.Println(err)
		
	}
	log.Logrecord("排程","schedule name: "+inventory.Name)
	ScreenshotbySchedule(scheduleID)
	// time.Sleep(3 * time.Second) 
	CreateHtmlbySchedule(scheduleID)
	CreatePDFbySchedule(scheduleID)
	SendEmailBySchedule(scheduleID)

}
