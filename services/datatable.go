package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)



// 新增element
func CreateDataTable(table entities.DataTable) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.DataTable{}
	err := global.Mysql.Create(&table).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Success = true
	res.Msg = "Create Success"
	// global.Mysql.Where("name = ?", element.Name).Omit("Instance").First(&res.Body)

	return res
}