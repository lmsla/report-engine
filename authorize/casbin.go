package authorize

import (
	"fmt"
	"report-backend-golang/global"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func LoadCasbin() {

	// client := env.Config.Database.Client
	host := global.EnvConfig.Database.Host
	port := global.EnvConfig.Database.Port
	user := global.EnvConfig.Database.User
	password := global.EnvConfig.Database.Password
	dbname := global.EnvConfig.Database.Db

	a, err := gormadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname), true)
	if err != nil {
		panic(err.Error())
	}
	global.CasbinEnforcer, err = casbin.NewEnforcer("./resources/rbac_model.conf", a)
	if err != nil {
		panic(err.Error())
	}
	err = global.CasbinEnforcer.LoadPolicy()
	if err != nil {
		panic(err.Error())
	}
}
