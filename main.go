package main

import (
	databases "report-backend-golang/database"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/router"
	"report-backend-golang/utils"
)

// @title Golang API - WarRoom
// @version 1.0
// @description Golang API 專案描述
// @termsOfService http://swagger.io/terms/

// @contact.name Jessie
// @contact.email support@swagger.io

// @host localhost:8005

// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @schemes http
func main() {

	utils.LoadEnvironment()

	databases.LoadDatabase()
	mysql, _ := global.Mysql.DB()
	defer mysql.Close()

	if global.EnvConfig.Other.Migration {
		entities.InitTable()
	}

	// utils.LoadCrontab()
	// authorize.LoadCasbin()

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)

}
