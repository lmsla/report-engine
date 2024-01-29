package services

import (
	// "fmt"
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/models"
	"time"
)

// 取得所有history
func GetAllHistory() models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.History{}

	err := global.Mysql.Find(&res.Body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}

	res.Success = true
	res.Msg = "Get All History Success"
	return res
}



// 新增history
func CreateHistory(history entities.History) models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []models.Instance{}

	err := global.Mysql.Create(&history).Error
	if err != nil {
		res.Msg = "Create Fail"
		return res
	}

	res.Success = true
	res.Msg = "Create Success"

	return res
}


// 取得符合刪除條件的 history
func GetOldHistory() models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.History{}

	// err := global.Mysql.Find(&res.Body).Error
	// if err != nil {
	// 	res.Msg = err.Error()
	// 	return res
	// }

	// 一个月前的时间戳
	oneMonthAgo := time.Now().AddDate(0, -1, 0).Unix()

	err := global.Mysql.Where("created_at < ?",oneMonthAgo).Find(&res.Body).Error
	if err != nil {
		res.Msg = err.Error()
		return res
	}
	// fmt.Println(oneMonthAgo)

	res.Success = true
	res.Msg = "Get All History Success"
	return res
}












func DeleteOldHistory() models.Response {

	res := models.Response{}
	res.Success = false
	res.Body = []entities.History{}

	// 一个月前的时间戳
	oneMonthAgo := time.Now().AddDate(0, -1, 0).Unix()

	//刪除 element 中相應的圖表
	err := global.Mysql.Where("created_at < ?",oneMonthAgo).Delete(&entities.History{}).Error
	if err != nil {
		res.Msg = fmt.Sprintf("Error when deleting Old History , err: %s", err)
		return res
	}

	res.Success = true
	res.Msg = "Delete Old History Success"

	return res

}
