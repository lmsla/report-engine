package main

import (
	// "fmt"
	"report-backend-golang/clients"
	"report-backend-golang/global"
	"report-backend-golang/migrate"
	"report-backend-golang/router"
	"report-backend-golang/utils"
	// "report-backend-golang/entities"
	// "report-backend-golang/log"
	// "report-backend-golang/tools"
	// "report-backend-golang/models"
	"report-backend-golang/services"

)

// @title Report Engine Golang API
// @version 1.0
// @description Golang API 專案描述
// @termsOfService http://swagger.io/terms/
// @contact.name Russell
// @contact.email support@swagger.io

//// @host 10.99.1.133:8005
// @host localhost:8005
// @BasePath  /api/v1
// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func main() {

	utils.LoadEnvironment()    
	// es.LoadElasticsearch()
	clients.LoadDatabase()
	mysql, _ := global.Mysql.DB()
	defer mysql.Close()

	// if global.EnvConfig.Database.Migration {
	// 	entities.InitTable()
	// }
	migrate.Run()

	utils.LoadCrontab()
	// // authorize.LoadCasbin()
	clients.LoadKeycloak()

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)

	// log.Logrecord("環境參數", "截圖檔案位置 : "+global.EnvConfig.Files.ScreenshotFile)
	// log.Logrecord("環境參數", "Html檔案位置 : "+global.EnvConfig.Files.HtmlFile)
	// log.Logrecord("環境參數", "PDF報表檔案位置 : "+global.EnvConfig.Files.ReportFile)
}

func main1() {

	// instance := models.Instance{
	// 	ID:       0,
	// 	Type:     "",
	// 	Name:     "",
	// 	URL:      "10.99.1.93:5601",
	// 	User:     "elastic",
	// 	Password: "12345678",
	// 	Auth:     0,
	// }
	utils.LoadEnvironment()
	// services.GetDataViewData("csc",instance,"bd9283f8-460b-48ab-bd23-f253bca38f12")

	clients.LoadDatabase()
	mysql, _ := global.Mysql.DB()
	defer mysql.Close()

	// services.DataDealing()
	// services.Testgy()
	services.Test1()
	// services.ExampleFpdf_CellFormat_tables1()
	// services.TableCreate()
	// dateString := "2024-09-25"

	// date,err := time.Parse("2006-01-02T15:04",dateString)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("date: %v\n", date)

}




