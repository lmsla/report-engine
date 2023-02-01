package services

import (
	// "fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
	"fmt"
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

	err := global.Mysql.Debug().Preload("Reports").Find(&body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}
	res.Body = body
	res.Success = true
	res.Msg = "Get All Schedule Success"
	return res
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

	// //先刪除 element 中相應的圖表
	// err := global.Mysql.Where("instance_id = ?", id).Delete(&entities.Element{}).Error
	// if err != nil {
	// 	res.Msg = fmt.Sprintf("Error when deleting related elements, err: %s", err)
	// 	return res
	// }

	err := global.Mysql.Where("id = ?", id).Delete(&entities.Schedule{}).Error
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
	res.Success = false
	res.Body = []entities.Schedule{}

	result := global.Mysql.Where("name = ?", schedule.Name).First(&entities.Schedule{})
	if result.RowsAffected > 0 {
		res.Msg = "Schedule Name already existed"
		return res
	}

	err := global.Mysql.Create(&schedule).Omit("Reports.Elements").Error
	if err != nil {
		res.Msg = "Create Fail"
		return res
	}

	res.Success = true
	res.Msg = "Create Success"
	global.Mysql.Where("name = ?", schedule.Name).First(&res.Body)

	return res
}