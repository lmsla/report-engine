package databases

import (
	"fmt"
	"time"
	"report-backend-golang/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadDatabase() {

	// 取得config參數
	host := global.EnvConfig.Database.Host
	port := global.EnvConfig.Database.Port
	user := global.EnvConfig.Database.User
	password := global.EnvConfig.Database.Password
	dbname := global.EnvConfig.Database.Db
	parameter := global.EnvConfig.Database.Params

	var err error
	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s\n", user, password, host, port, dbname, parameter)
	global.Mysql, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbname, parameter)), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := global.Mysql.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(int(global.EnvConfig.Database.MaxIdle))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(int(global.EnvConfig.Database.MaxOpenConn))

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	lifeTime, _ := time.ParseDuration(global.EnvConfig.Database.MaxLifeTime)
	sqlDB.SetConnMaxLifetime(lifeTime)
	fmt.Println("Database 連線成功")
}
