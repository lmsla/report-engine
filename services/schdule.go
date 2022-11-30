package services

import (
	// "fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

func GetAllSchedule() models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.Schedule{}

	err := global.Mysql.Find(&res.Body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = "Get All Schedule Success"
	return res
}



// 新增schedule
func CreateSchedule(schedule entities.Schedule) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.Schedule{}

	result := global.Mysql.Where("name = ?", schedule.Name).First(&entities.Schedule{})
	if result.RowsAffected > 0 {
		res.Msg = "Instance Name already existed"
		return res
	}

	err := global.Mysql.Create(&schedule).Error
	if err != nil {
		res.Msg = "Create Fail"
		return res
	}

	res.Success = true
	res.Msg = "Create Success"
	global.Mysql.Where("name = ?", schedule.Name).First(&res.Body)

	return res
}