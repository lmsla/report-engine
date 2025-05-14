package services

import (
	// "fmt"
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

func GetAllSchedule() models.Response {

	res := models.Response{}
	res.Success = false
	var body = []entities.Schedule{}

	// err := global.Mysql.Find(&body).Error
	// if err != nil {
	// 	res.Msg = err.Error()
	// 	return res
	// }

	err := global.Mysql.Preload("Reports").Find(&body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}
	res.Body = body
	res.Success = true
	res.Msg = "Get All Schedule Success"
	return res
}

// 查單一Schedule by schedule ID
func GetScheduleBySheduleID(scheduleID int) (entities.Schedule, error) {

	var instance entities.Schedule
	instance.ID = scheduleID
	err := global.Mysql.Preload("Reports").First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}

func DeleteSchedule(id int) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = nil

	result := global.Mysql.Where("id = ?", id).First(&entities.Schedule{})
	if result.RowsAffected == 0 {
		res.Msg = "Schedule ID does not exist"
		return res
	}

	//先刪除 ReportsSchedule 中相應的圖表
	err := global.Mysql.Where("schedule_id = ?", id).Delete(&entities.ReportsSchedules{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting data in ReportsSchedule, err: %s", err)
		return res
	}
	data, err := GetEntryByScheduleID(id)
	if err != nil {
		res.Msg = fmt.Sprintf("Error when get data in CronList, err: %s", err)
		return res
	}

	// 找出目前 Entries 中 與CronList 對應的 entryID 並刪除 Entries 中的 entry
	entries := global.Crontab.Entries()
	for _, entry := range entries {
		if data.EntryID == int(entry.ID) {
			global.Crontab.Remove(entry.ID)
		}
	}

	// 刪除 CronList 中對應的 entry ID
	err = global.Mysql.Where("schedule_id = ?", id).Delete(&entities.CronList{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting data in CronList, err: %s", err)
		return res
	}

	err = global.Mysql.Where("id = ?", id).Delete(&entities.Schedule{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting Schedule, err: %s", err)
		return res
	}

	res.Success = true
	res.Msg = "Delete Success"

	return res

}

// 新增schedule
func CreateSchedule(schedule entities.Schedule) models.Response {

	res := models.Response{}
	// res.Success = false
	// res.Body = []entities.Schedule{}

	result := global.Mysql.Where("name = ?", schedule.Name).First(&entities.Schedule{})
	if result.RowsAffected > 0 {
		res.Msg = "Schedule Name already existed"
		return res
	}

	err := global.Mysql.Create(&schedule).Omit("Reports.Elements").Error
	if err != nil {
		res.Msg = "Create Fail"
		return res
	} else {
		fmt.Println("schedule.ID")
		fmt.Println(schedule.ID)
		ExecuteShedulePDF(schedule.ID)
	}

	res.Msg = "Create Success"
	res.Success = true
	// global.Mysql.Where("name = ?", schedule.Name).First(&res.Body)
	return res
}

func UpdateSchedule(schedule entities.Schedule) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.Schedule{}

	// result := global.Mysql.Where("id != ? AND name = ?", instance.ID, instance.Name).First(&entities.Instance{})
	// if result.RowsAffected > 0 {
	// 	res.Msg = "Instance Name already existed"
	// 	return res
	// }

	//先刪除 ReportsSchedule 中相應的圖表
	err := global.Mysql.Where("schedule_id = ?", schedule.ID).Delete(&entities.ReportsSchedules{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting data in ReportsSchedule, err: %s", err)
		return res
	}
	data, err := GetEntryByScheduleID(schedule.ID)
	if err != nil {
		res.Msg = fmt.Sprintf("Error when get data in CronList, err: %s", err)
		return res
	}

	// 找出目前 Entries 中 與CronList 對應的 entryID 並刪除 Entries 中的 entry
	entries := global.Crontab.Entries()
	for _, entry := range entries {
		if data.EntryID == int(entry.ID) {
			global.Crontab.Remove(entry.ID)
		}
	}

	// 刪除 CronList 中對應的 entry ID
	err = global.Mysql.Where("schedule_id = ?", schedule.ID).Delete(&entities.CronList{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting data in CronList, err: %s", err)
		return res
	}

	//update schedule
	err = global.Mysql.Select("*").Where("id = ?", schedule.ID).Updates(&schedule).Error
	if err != nil {
		res.Msg = "Update Fail"
		return res
	} else {
		res.Msg = "Update Success"
		// enable = true 時才啟用排程
		if schedule.Enable {
			ExecuteShedulePDF(schedule.ID)
		}

	}

	res.Success = true
	res.Msg = "Update Success"
	global.Mysql.Where("id = ?", schedule.ID).First(&res.Body)

	return res

}

func GetEntryByScheduleID(scheduleID int) (entities.CronList, error) {
	cronlist := entities.CronList{}

	err := global.Mysql.Debug().Where("schedule_id = ?", scheduleID).Find(&cronlist).Error
	if err != nil {
		return cronlist, err
	}
	return cronlist, nil

}
