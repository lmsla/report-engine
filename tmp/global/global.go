package global

import (
	"report-backend-golang/structs"

	"github.com/gomodule/redigo/redis"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var (
	EnvConfig *structs.EnviromentModel
	Mysql     *gorm.DB
	Redis     *redis.Pool
	Crontab   *cron.Cron
)
