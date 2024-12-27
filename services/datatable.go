package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

func GetTables() models.Response {

	res := models.Response{}
	res.Success = false
	var body = []entities.Table{}
	err := global.Mysql.Debug().Preload("Instance").Preload("Columns").Find(&body).Error
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
func CreateTable(table entities.Table) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.Table{}
	err := global.Mysql.Create(&table).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Success = true
	res.Msg = "Create Success"
	global.Mysql.Where("name = ?", table.Name).Omit("Report").Omit("Instance").First(&res.Body)

	return res
}

func UpdateTable(table entities.Table) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.Table{}

	//先刪除 column 中相應的欄位
	err := global.Mysql.Where("table_id = ?", table.ID).Delete(&entities.Column{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting related columns, err: %s", err)
		return res
	}

	err = global.Mysql.Select("*").Where("id = ?", table.ID).Updates(&table).Error
	if err != nil {
		res.Msg = "Update Fail"
		return res
	}

	res.Success = true
	res.Msg = "Update Success"
	global.Mysql.Where("id = ?", table.ID).First(&res.Body)

	return res

}

func DeleteTable(id int) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = nil

	result := global.Mysql.Where("id = ?", id).First(&entities.Table{})
	if result.RowsAffected == 0 {
		res.Msg = "DataTable ID does not exist"
		return res
	}

	err := global.Mysql.Where("id = ?", id).Delete(&entities.Table{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting DataTable, err: %s", err)
		return res
	}

	res.Success = true
	res.Msg = "Delete Success"

	return res

}

// 查單一Table
func GetTableByID(tableID int) (entities.Table, error) {

	var table entities.Table
	table.ID = tableID
	err := global.Mysql.Debug().Preload("Instance").Preload("Columns").First(&table).Error
	if err != nil {
		return table, err
	}
	return table, nil
}

// 查 Tables by ReportID
func GetTableByReportID(reportID int) ([]entities.Table, error) {
	// element :=  entities.Element{}
	report := entities.Report{}
	err := global.Mysql.Debug().Where("id = ?", reportID).Preload("Tables").Preload("Tables.Columns").Preload("Tables.Instance").Find(&report).Error
	if err != nil {
		return nil, err
	}
	return report.Tables, nil
}
