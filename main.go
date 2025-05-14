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
	// "report-backend-golang/services"

)

// @title Report Engine Golang API
// @version 1.0
// @description Golang API 專案描述
// @termsOfService http://swagger.io/terms/
// @contact.name Russell
// @contact.email support@swagger.io
// @host 10.99.1.213:8005
//// @host localhost:8005
// @BasePath  /api/v1
// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http
func main() {

	utils.LoadEnvironment()    
	
	clients.LoadDatabase()
	mysql, _ := global.Mysql.DB()
	defer mysql.Close()

	migrate.Run()

	utils.LoadCrontab()

	clients.LoadKeycloak()

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)

	// log.Logrecord("環境參數", "截圖檔案位置 : "+global.EnvConfig.Files.ScreenshotFile)
	// log.Logrecord("環境參數", "Html檔案位置 : "+global.EnvConfig.Files.HtmlFile)
	// log.Logrecord("環境參數", "PDF報表檔案位置 : "+global.EnvConfig.Files.ReportFile)
}




