package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/log"
	"time"
)

// 執行 PDF Schedule by ScheduleID
func ExecuteShedulePDF(scheduleID int) {

	inventory, err := GetScheduleBySheduleID(scheduleID)
	if err != nil {
		fmt.Println(err)
	}
	history := entities.History{}
	history.ScheduleID = scheduleID
	// history.To = inventory.To
	fmt.Println("ScheduleID in history: ")
	fmt.Println(history.ScheduleID)

	// inventory.CronID
	EntryID, err := global.Crontab.AddFunc(inventory.CronTime, func() {
		fmt.Println("執行排程")
		test := false
		FuncAddToCron(test, scheduleID)
	})
	fmt.Println("entryID: ")
	fmt.Println(EntryID, err)

	// 寫一筆記錄到 cron_lists 的 table 中
	// res := models.Response{}
	// res.Success = false
	cronlist := entities.CronList{ScheduleID: scheduleID, EntryID: int(EntryID)}
	result := global.Mysql.Create(&cronlist).Error
	if result != nil {
		fmt.Println("cronlist Create Fail")
		// return res)
	}

	//fmt.Print(global.EnvConfig.CRONTAB.Period,global.EnvConfig.INFLUX.URL)
	if err != nil {
		fmt.Println("crontab PDF 初始化失敗")
		// log.Logrecord("排程 ","PDF排程 初始化失敗")
		fmt.Println(err.Error())
		log.Logrecord("ERROR", "PDF排程 初始化失敗"+err.Error())
	} else {

		fmt.Println("crontab PDF 初始化成功")
		log.Logrecord("排程", "PDF排程 初始化成功")
		global.Crontab.Start()

	}

}

func FuncAddToCron(test bool, scheduleID int) {
	inventory, err := GetScheduleBySheduleID(scheduleID)
	if err != nil {
		log.Logrecord("ERROR", "Get Schedule by Schedule Id error"+err.Error())
		fmt.Println(err)

	}
	time_execute := time.Now().Unix()
	fmt.Println("排程執行時間: ", time_execute)
	history := entities.History{}
	history.ScheduleID = scheduleID
	history.To = inventory.To
	history.BCC = inventory.BCC
	history.CC = inventory.CC
	history.ScheduleName = inventory.Name
	history.ExecuteTime = time_execute
	if test {
		history.Success = "試寄成功"
	} else {
		history.Success = "成功"
	}
	log.Logrecord("排程", "schedule name: "+inventory.Name)
	history.Success = "成功"
	// 移除掉 CreateHtmlbySchedule(time_execute, scheduleID)
	err = ScreenshotbySchedule(time_execute, scheduleID)
	if err != nil {
		fmt.Println("截圖錯誤", err)
		history.Success = "截圖錯誤"
	} else {
		// CreatePDFbySchedule(scheduleID)
		err = CreatePDFbySchedule(time_execute, scheduleID)
		if err != nil {
			fmt.Println("產生pdf錯誤", err)
			history.Success = "產生pdf錯誤"
		} else {
			err = SendEmailBySchedule(time_execute, scheduleID)
			if err != nil {
				fmt.Println("發送mail錯誤", err)
				history.Success = "發送mail錯誤"
			}

		}

	}

	time_mail := time.Now().Unix()
	fmt.Println("寄送時間:", time_mail)
	history.EmailTime = time_mail

	err = global.Mysql.Create(&history).Error
	if err != nil {
		msg := fmt.Sprintf("create history error: %s", err.Error())
		log.Logrecord("ERROR", msg)
	}
	if test {
		log.Logrecord("排程", fmt.Sprintf("報表試寄完成, Schedule ID : %d", scheduleID))
	}
	

	DeleteOldHistory()
}

// report 重寄by 歷史紀錄
func CreateHistoryReport(historyID int) {

	historydata, err := GetHistoryByHistoryID(historyID)
	if err != nil {
		log.Logrecord("ERROR", "Get History by History Id error"+err.Error())
		// fmt.Println(err)
	}

	inventory, err := GetScheduleBySheduleID(historydata.ScheduleID)
	if err != nil {
		log.Logrecord("ERROR", "Get Schedule by Schedule Id error"+err.Error())
		// fmt.Println(err)

	}

	// 以歷史執行時間為新的執行時間
	history_ExecuteTime := historydata.ExecuteTime

	// 新的歷史執行時間
	time_execute := time.Now().Unix()
	fmt.Println("排程執行時間: ", time_execute)
	history := entities.History{}
	history.ScheduleID = historydata.ScheduleID
	history.To = inventory.To
	history.BCC = inventory.BCC
	history.CC = inventory.CC
	history.ScheduleName = inventory.Name
	history.ExecuteTime = time_execute

	log.Logrecord("排程", "schedule name: "+inventory.Name)
	history.Success = "成功"

	err = ScreenshotbySchedule(history_ExecuteTime, historydata.ScheduleID)
	fmt.Println("ScreenshotbySchedule err", err)
	if err != nil {
		fmt.Println("截圖錯誤", err)
		history.Success = "截圖錯誤"
	} else {
		err = CreatePDFbySchedule(history_ExecuteTime, historydata.ScheduleID)
		if err != nil {
			fmt.Println("產生pdf錯誤", err)
			history.Success = "產生pdf錯誤"
		} else {
			err = SendEmailBySchedule(history_ExecuteTime, historydata.ScheduleID)
			if err != nil {
				fmt.Println("發送mail錯誤", err)
				history.Success = "發送mail錯誤"
			}

		}

	}
	// time_mail := time.Now().Format("2006-01-02 15:04:05")
	//寄送時間
	time_mail := time.Now().Unix()
	history.EmailTime = time_mail

	err = global.Mysql.Create(&history).Error
	if err != nil {
		msg := fmt.Sprintf("create history error: %s", err.Error())
		log.Logrecord("ERROR", msg)
	}

}
