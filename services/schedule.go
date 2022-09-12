package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
	// "report-backend-golang/screenshot"
	// "report-backend-golang/schedule"
	// "github.com/robfig/cron/v3"
	// "report-backend-golang/log"
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
		// schedule.ExecuteSheduleSendMail(schedule.ScheduleID)
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
