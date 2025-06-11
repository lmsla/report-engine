package services

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

func GetAllElements() models.Response {

	res := models.Response{}
	res.Success = false
	// res.Body = []entities.Element{}
	var body = []entities.Element{}
	err := global.Mysql.Preload("Instance").Find(&body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}

	res.Body = body
	res.Success = true
	res.Msg = "Get All Elements Success"
	return res
}


// 查 Elements by ReportID
func GetElementsByReportID(reportID int) ([]entities.Element,error) {

	// element :=  entities.Element{}
	report := entities.Report{}
	// report.ID = reportID
	// res := models.Response{}
	// res.Success = false
	// var body = []entities.Element{}
	err := global.Mysql.Debug().Where("id = ?",reportID).Preload("Elements").Preload("Elements.Instance").Find(&report).Error
	if err != nil {
		return nil,err
	}
	return report.Elements,nil
}

// // 查單一 Element by ReportID
// func GetElementsByReportID(reportID int) (models.Response) {

// 	// element :=  entities.Element{}

// 	res := models.Response{}
// 	res.Success = false
// 	var body = []entities.Element{}
// 	err := global.Mysql.Debug().Where("report_id = ?",reportID).Preload("Instance").Find(&body).Error
// 	if err != nil {
// 		res.Msg = err.Error()
// 		return res
// 	}

// 	res.Body = body
// 	res.Success = true
// 	res.Msg = "Get Selected Report Success"
// 	return res


// }


// 新增element
func CreateElement(element entities.Element) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.Element{}
	err := global.Mysql.Create(&element).Error

	if err != nil {
		res.Msg = fmt.Sprintf("Error: %v", err)
		res.Success = false
		return res
	}
	res.Success = true
	res.Msg = "Create Elements Success"
	global.Mysql.Where("name = ?", element.Name).Omit("Instance").First(&res.Body)

	return res
}


func UpdateElement(element entities.Element) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.Element{}

	// result := global.Mysql.Where("id != ? AND name = ?", instance.ID, instance.Name).First(&entities.Instance{})
	// if result.RowsAffected > 0 {
	// 	res.Msg = "Instance Name already existed"
	// 	return res
	// }

	err := global.Mysql.Select("*").Where("id = ?", element.ID).Updates(&element).Error
	if err != nil {
		res.Msg = "Update Elements Fail"
		return res
	}

	res.Success = true
	res.Msg = "Update Elements Success"
	global.Mysql.Where("id = ?", element.ID).First(&res.Body)

	return res

}




func DeleteElement(id int) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = nil

	result := global.Mysql.Where("id = ?", id).First(&entities.Element{})
	if result.RowsAffected == 0 {
		res.Msg = "Element ID does not exist"
		return res
	}

	// //先刪除 element 中相應的圖表
	// err := global.Mysql.Where("instance_id = ?", id).Delete(&entities.Element{}).Error
	// if err != nil {
	// 	res.Msg = fmt.Sprintf("Error when deleting related elements, err: %s", err)
	// 	return res
	// }

	err := global.Mysql.Where("id = ?", id).Delete(&entities.Element{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting element, err: %s", err)
		return res
	}

	res.Success = true
	res.Msg = "Delete Success"

	return res

}