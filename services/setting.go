package services

import (
	"fmt"
	"report-backend-golang/global"
	"report-backend-golang/entities"
	"gorm.io/gorm/clause"
)


func GetServerModule() ([]entities.Module,error ){
	db := global.Mysql
	modules := []entities.Module{}
	err := db.Model(entities.Module{}).Where(&entities.Module{
		Disabled: false}).Find(&modules).Error
	if err != nil {
		// global.Logger.Error(
		// 	err.Error(),
		// 	zap.String(global.LogEvent.Tag.Service, global.LogEvent.API.Enviroment),
		// )
		fmt.Println("GetServerModule error", err.Error())
		return nil,err
	}
	return modules,err
}

func GetServerMenu() ([]entities.MainMenu, error) {
	db := global.Mysql
	menus := []entities.MainMenu{}
	// var err error

	err := db.Model(entities.MainMenu{}).Preload(clause.Associations).Find(&menus).Error

	// switch role {
	// case global.EnvConfig.SSO.AdminRole:

	// 	err = db.Model(entities.MainMenu{}).Preload(clause.Associations).Find(&menus).Error
	// default:

	// 	err = db.Model(entities.MainMenu{}).Preload(clause.Associations).Not(
	// 		&entities.MainMenu{
	// 			OnlyAdmin: true}).Find(&menus).Error
	// }
	if err != nil {
		
		// global.Logger.Error(
		// 	err.Error(),
		// 	zap.String(global.LogEvent.Tag.Service, global.LogEvent.API.Enviroment),
		// )
		fmt.Println("GetServerMenu error", err.Error())
		return nil, err
	}
	for _, data := range menus {
		fmt.Println(data.Icon)
	}
	fmt.Println(menus)
	return menus, nil
}
