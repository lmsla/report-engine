package migrate

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"report-backend-golang/global"
)

func Run() {
	// 取得config參數
	host := global.EnvConfig.Database.Host
	port := global.EnvConfig.Database.Port
	user := global.EnvConfig.Database.User
	password := global.EnvConfig.Database.Password
	dbname := global.EnvConfig.Database.Db
	parameter := global.EnvConfig.Database.Params

	// 建立資料庫連線
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, dbname, parameter))
	if err != nil {
		fmt.Printf("[ERROR] Database connection failed: %v\n", err)
		panic(err)
	}
	defer db.Close()

	fmt.Println("[SUCCESS] Migration database connection success!")

	// 建立 MySQL driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Printf("[ERROR] MySQL driver initialization failed: %v\n", err)
		panic(err)
	}

	fmt.Println("[SUCCESS] MySQL driver initialized")

	// 建立 migration instance
	m, err := migrate.NewWithDatabaseInstance(
		global.EnvConfig.Files.MigrationsFile,
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Printf("[ERROR] Migration instance creation failed: %v\n", err)
		panic(err)
	}

	fmt.Println("[SUCCESS] Migration instance ready")

	// 檢查當前 migration 狀態
	if version, dirty, err := m.Version(); err == nil {
		fmt.Printf("[INFO] Current migration version: %d (dirty: %v)\n", version, dirty)
		if dirty {
			fmt.Println("[WARNING]: Database is in dirty state, migration may have failed previously")
		}
	} else {
		fmt.Println("[INFO] No previous migrations found, starting fresh migration")
	}

	// 執行 migration
	fmt.Println("[PROCESS] Starting migration process...")
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("[SUCCESS] Database schema is up to date, no migrations to apply")
		} else {
			fmt.Printf("[ERROR] Migration failed: %v\n", err)
			panic(err)
		}
	} else {
		fmt.Println("[SUCCESS] Migration completed successfully")
	}

	// 顯示最終狀態
	if version, dirty, err := m.Version(); err == nil {
		fmt.Printf("[INFO] Migration completed. Current version: %d (dirty: %v)\n", version, dirty)
	}

	fmt.Println("[COMPLETED] Migration process finished")
}
