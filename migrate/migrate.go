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

	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbname, parameter))
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"mysql",
		driver,
	)

	// // var err error
	// err := errors.New("mock error")
	// for err != nil {
	// 	global.Mysql, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbname, parameter)), &gorm.Config{
	// 		DisableForeignKeyConstraintWhenMigrating: true,
	// 	})
	// 	time.Sleep(1 * time.Second)
	// }
	fmt.Println(m.Up())
	// or m.Step(2) if you want to explicitly set the number of migrations to run
}
