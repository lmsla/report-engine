package entities

import (
	"os"
	"report-backend-golang/global"

	"github.com/gookit/color"
)

func InitTable() {
	var err error

	err = global.Mysql.Migrator().DropTable(
		Instance{},
		Schedule{},
		Report{},
		Element{},
		History{},
		"reports_schedules",
	)
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化數據失敗,err: %v\n", err)
		os.Exit(0)
	}

	err = global.Mysql.AutoMigrate(
		Instance{},
		Schedule{},
		Report{},
		Element{},
		History{},
	)

	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化數據失敗,err: %v\n", err)
		os.Exit(0)
	}

	err = global.Mysql.Create(&InstanceData).Error
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化數據失敗,err: %v\n", err)
		os.Exit(0)
	}

	err = global.Mysql.Create(&ScheduleData).Error
	if err != nil {
		color.Warn.Printf("[Mysql]-->初始化數據失敗,err: %v\n", err)
		os.Exit(0)
	}

	global.Mysql.Migrator().CreateConstraint(&Report{}, "Element")
	global.Mysql.Migrator().CreateConstraint(&Element{}, "Instance")
	global.Mysql.Migrator().CreateConstraint(&Report{}, "Schedules")
	global.Mysql.Migrator().CreateConstraint(&Schedule{}, "Reports")
	global.Mysql.Migrator().CreateConstraint(&History{}, "Schedule")
	global.Mysql.Migrator().CreateConstraint(&History{}, "Report")

	color.Info.Println("[Mysql]-->初始化數據成功")
}

var InstanceData = []Instance{
	{
		Type:     "kibana",
		Name:     "kibana_110",
		URL:      "http://10.99.1.110:5601/kibana_iframe",
		User:     "elasic",
		Password: "RnIv7YhigaVKS=l-*yz9",
		Auth:     0,
	},
	{
		Type:     "grafana",
		Name:     "grafana_241",
		URL:      "http://10.99.1.241:3000",
		User:     "admin",
		Password: "12345678",
		Auth:     0,
	},
}

var ReportData = []Report{
	{
		Name:       "Kibana_Test",
		TimeUnit:   "日",
		TimePeriod: 1,
		Elements: []Element{
			{
				Type:       "dashboard",
				Name:       "cht_vsm_MD_track_dashboard",
				UID:        "cdd23320-edf8-11ec-b4a6-b712e673cd3e",
				RowNum:     1,
				ColumnType: "M",
				InstanceID: 1,
				SpaceName: "default",
			},
			{
				Type:       "visualization",
				Name:       "cht_apache_link",
				UID:        "c16aa400-ef01-11ec-b4a6-b712e673cd3e",
				RowNum:     3,
				ColumnType: "L",
				InstanceID: 1,
				SpaceName: "default",
			},
			{
				Type:       "search",
				Name:       "cht_mysql_apache_discover",
				UID:        "60d92120-ee11-11ec-b4a6-b712e673cd3e",
				RowNum:     2,
				ColumnType: "M",
				InstanceID: 1,
				SpaceName: "default",
			},
		},
	},
	{
		Name:       "Grafana_Test",
		TimeUnit:   "日",
		TimePeriod: 1,
		Elements: []Element{
			{
				Type:       "dashboard",
				Name:       "Rack Monitor (Current)",
				UID:        "m-OiCeDVz",
				RowNum:     1,
				ColumnType: "M",
				InstanceID: 2,
				SpaceName: "default",
			},
		},
	},
}

var ScheduleData = []Schedule{
	{
		Name:     "日報",
		CronTime: "0 0 * * *",
		To:       []string{""},
		CC:       []string{""},
		BCC:      []string{""},
		CronID:   0,
		Reports:  ReportData,
	},
}
