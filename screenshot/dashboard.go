package screenshot

import (
	"fmt"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	// "report-backend-golang/models"
)

// // 查 instance 的資訊-backup

// func GetInstanceInfoByInstanceID(id int) (models.Response ){
// 	res :=  models.Response{}

// 	err := global.Mysql.Where("instance_id= ?",id).First(&entities.Instance{}).Error
// 	if err != nil {
// 		res.Msg = fmt.Sprintf("Error when check instance ID, error: %s",err.Error())
// 		res.Success = false
// 		return res
// 	}

// 	res.Msg = "ok"
// 	res.Success = true
// 	return res

// }


// 查 instance 的資訊
func GetInstanceInfoByInstanceID(InstanceID int) (entities.Instance, error){

	var instance entities.Instance
	instance.InstanceID = InstanceID

	err := global.Mysql.Where("instance_id= ?",InstanceID).First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil

}


// 查單一Instance
func GetInstanceByID(instanceID int) (entities.Instance, error) {

	var instance entities.Instance
	instance.InstanceID = instanceID
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}


// 只取出 Instance 中需要的元素
func GetInstanceTypeAndIP(id int) (map[string]string, error){
	// var IPType []map[string]string
	data,err := GetInstanceByID(id)
	if err != nil {
		fmt.Println(err)
		// log.Logrecord("ERROR",err.Error())
		//return
	}
	// var UID1 []string
	IPType := make(map[string]string)
	IPType["IP"] = data.IP
	IPType["type"] = data.Type
	IPType["user"] = data.User
	IPType["password"] = data.Pass
	return IPType,err
}
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

// 查詢All dashboard
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