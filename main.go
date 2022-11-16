package main

import (
	"report-backend-golang/clients"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/router"
	"report-backend-golang/utils"
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

	if global.EnvConfig.Database.Migration {
		entities.InitTable()
	}

	clients.LoadRedis()
	defer global.Redis.Close()

	utils.LoadCrontab()
	// // authorize.LoadCasbin()

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)

	// 測試 rule
	// services.CheckRule(2)

}
