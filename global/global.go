package global

import (
	"github.com/robfig/cron/v3"
	"report-backend-golang/structs"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

var (
	EnvConfig      *structs.EnviromentModel
	Mysql          *gorm.DB
	Session        *gorm.Session
	CasbinEnforcer *casbin.Enforcer
	Crontab        *cron.Cron
)
