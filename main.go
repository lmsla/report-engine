package main

import (
	"report-backend-golang/clients"
	"report-backend-golang/global"
	"report-backend-golang/migrate"
	"report-backend-golang/router"
	"report-backend-golang/services"
	// "report-backend-golang/log"
	// "report-backend-golang/tools"
	// "report-backend-golang/services"
	"report-backend-golang/utils"
	// "report-backend-golang/entities"
)

// @title Report Engine Golang API
// @version 1.0
// @description Golang API 專案描述
// @termsOfService http://swagger.io/terms/

// @contact.name Winston
// @contact.email support@swagger.io

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

	clients.LoadRedis()
	defer global.Redis.Close()

	utils.LoadCrontab()
	// // authorize.LoadCasbin()

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)

	// log.Logrecord("環境參數", "截圖檔案位置 : "+global.EnvConfig.Files.ScreenshotFile)
	// log.Logrecord("環境參數", "Html檔案位置 : "+global.EnvConfig.Files.HtmlFile)
	// log.Logrecord("環境參數", "PDF報表檔案位置 : "+global.EnvConfig.Files.ReportFile)
}

func main1() {
	utils.LoadEnvironment()
	utils.LoadCrontab()
	// // authorize.LoadCasbin()
	// services.SendEmail()
	// tools.Timeconverter("年",1)
	services.CronTest()

}
