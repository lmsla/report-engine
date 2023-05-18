package services

import (
	// "fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)


// 新增instance
func CreateHistory(history entities.History) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []models.Instance{}

	// result := global.Mysql.Where("name = ?", instance.Name).First(&entities.Instance{})
	// if result.RowsAffected > 0 {
	// 	res.Msg = "Instance Name already existed"
	// 	return res
	// }

	err := global.Mysql.Create(&history).Error
	if err != nil {
		res.Msg = "Create Fail"
		return res
	}

	res.Success = true
	res.Msg = "Create Success"
	// global.Mysql.Where("name = ?", instance.Name).First(&res.Body)

	return res
}
