package screenshot

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	// "report-backend-golang/models"
)



// 查 Dashboard uid by report ID
func GetDashboardByReportID(reportID int) (entities.Report, error) {

	var instance entities.Report
	instance.ReportID = reportID
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}

// 查詢All group
func GetAllDashboard() ([]entities.RDashboard, error) {

	var instancies []entities.RDashboard
	err := global.Mysql.Find(&instancies).Error
	if err != nil {
		return nil, err
	}

	return instancies, nil
}

func GetUID() ([]string,error){
	uid,err := GetAllDashboard()
	if err != nil {
		fmt.Println(err)
		// log.Logrecord("ERROR",err.Error())
		//return
	}
	var UID1 []string
	for _,uids := range uid {
		// var UID1 []string
		UID1 = append(UID1,uids.UID)
	}	
	fmt.Println(UID1)
	if err != nil {
		return nil, err
	}
	return UID1,nil

}


// fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s\n", user, password, host, port, dbname, parameter)