package crontab

import (
	"github.com/robfig/cron/v3"
	// "log"
	// "report-backend-golang/global"
	"report-backend-golang/log"

	"fmt"
)




func LoadCrontab(Period string ) {
	//checkLIcense()
	c := cron.New()
	_, err := c.AddFunc(Period, test)
	//fmt.Print(global.EnvConfig.CRONTAB.Period,global.EnvConfig.INFLUX.URL)
	if err != nil {
		fmt.Println("crontab xdr 初始化失敗")
		log.Logrecord("排程 ","xdr排程 初始化失敗")
		fmt.Println(err.Error())
		log.Logrecord("ERROR ",err.Error())
	} else {
		fmt.Println("crontab xdr 初始化成功")
		log.Logrecord("排程 ","xdr排程 初始化成功")
		c.Start()

	}
}

func test(){

}