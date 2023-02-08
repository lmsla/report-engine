package entities

type Schedule struct {
	Common
	ID       int      `gorm:"primaryKey;index" json:"id" form:"id"`
	Name     string   `gorm:"type:varchar(50)" json:"name" form:"name"`
	CronTime string   `gorm:"type:varchar(50)" json:"cron_time" form:"cron_time"`
	To       string   `gorm:"type:varchar(50)" json:"to" form:"to"`
	CC       string   `gorm:"type:varchar(50)" json:"cc" form:"cc"`
	BCC      string   `gorm:"type:varchar(50)" json:"bcc" form:"bcc"`
	CronID   int      `gorm:"type:int" json:"cron_id" form:"cron_id"`
	Reports  []*Report `gorm:"many2many:reports_schedules;foreignKey:ID;reference:ID;"`
}


// foreignKey:InstanceID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;


type ReportsSchedules struct {
	ReportID  int `gorm:"primaryKey" form:"report_id"`
	ScheduleID int `gorm:"primaryKey" form:"schedule_id"`
  }