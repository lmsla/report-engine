package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)



func GetDataTables() models.Response {

	res := models.Response{}
	res.Success = false
	var body = []entities.DataTable{}
	err := global.Mysql.Debug().Preload("Instance").Find(&body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}
	res.Body = body
	res.Success = true
	res.Msg = "Get All DataTable Success"
	return res
}



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
	global.Mysql.Where("name = ?", table.Name).Omit("Instance").First(&res.Body)

	return res
}


func UpdateDataTable(table entities.DataTable) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.DataTable{}

	err := global.Mysql.Select("*").Where("id = ?", table.ID).Updates(&table).Error
	if err != nil {
		res.Msg = "Update Fail"
		return res
	}

	res.Success = true
	res.Msg = "Update Success"
	global.Mysql.Where("id = ?", table.ID).First(&res.Body)

	return res

}



func DeleteDataTable(id int) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = nil

	result := global.Mysql.Where("id = ?", id).First(&entities.DataTable{})
	if result.RowsAffected == 0 {
		res.Msg = "DataTable ID does not exist"
		return res
	}

	err := global.Mysql.Where("id = ?", id).Delete(&entities.DataTable{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting DataTable, err: %s", err)
		return res
	}

	res.Success = true
	res.Msg = "Delete Success"

	return res

}