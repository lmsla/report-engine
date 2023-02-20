package services

import (
	"fmt"
	// "report-backend-golang/entities"
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
	// inventory1,err := services.GetReportByReportName(inventory.Report)

	
	_,err = global.Crontab.AddFunc(inventory.CronTime,func(){
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






func FuncAddToCron(scheduleID int) {
	inventory,err := GetScheduleBysSheduleID(scheduleID)
	if err != nil {
		fmt.Println(err)
	}
	log.Logrecord("排程","schedule name: "+inventory.Name)
	ScreenshotbySchedule(scheduleID)
	// time.Sleep(3 * time.Second) 
	CreateHtmlbySchedule(scheduleID)
	CreatePDFbySchedule(scheduleID)
	SendEmailBySchedule(scheduleID)

	// inventory1,err := GetReportByScheduleID(scheduleID)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _,data := range inventory1 {
	// 	log.Logrecord("排程","report name: "+data.Name + "開始執行截圖")
	// 	ScreenshotbyReport(data.ID)
	// 	log.Logrecord("排程","report name: "+data.Name + "開始產出報表")
	// 	CreateHtml(data.ID)
	// } 

	// Sendmail(scheduleID,inventory.Recipient)

}
