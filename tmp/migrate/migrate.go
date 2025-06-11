package migrate

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// "report-backend-golang/clients"
	"fmt"
	"report-backend-golang/global"
	// "gorm.io/driver/mysql"
)

func Run() {
	// 取得config參數
	host := global.EnvConfig.Database.Host
	port := global.EnvConfig.Database.Port
	user := global.EnvConfig.Database.User
	password := global.EnvConfig.Database.Password
	dbname := global.EnvConfig.Database.Db
	parameter := global.EnvConfig.Database.Params

	if db, e := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbname, parameter)); e != nil {
		fmt.Printf(e.Error())
		panic(e)
	} else {
		fmt.Println("migration db connect success!!")
		if driver, e := mysql.WithInstance(db, &mysql.Config{}); e != nil {
			fmt.Printf(e.Error())
			panic(e)
		} else {
			fmt.Println("driver ok")
			if m, e := migrate.NewWithDatabaseInstance(
				global.EnvConfig.Files.MigrationsFile,
				// "file://./db/migrations",
				// "file:///src/db/migrations"
				"mysql",
				driver,
			); e != nil {
				fmt.Printf(e.Error())
				panic(e)
			} else {
				fmt.Println("migration ready to go")
				fmt.Println(m.Up())
			}
		}
	}
	// or m.Step(2) if you want to explicitly set the number of migrations to run
}
