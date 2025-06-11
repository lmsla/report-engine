package utils

import (
	"report-backend-golang/global"

	"github.com/robfig/cron/v3"
)

func LoadCrontab() {

	global.Crontab = cron.New()
	global.Crontab.Start()
}
