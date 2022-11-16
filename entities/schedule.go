package entities

type Schedule struct {
	Common
	ID       int      `gorm:"primaryKey;index"`
	Name     string   `gorm:"type:varchar(50)"`
	CronTime string   `gorm:"type:varchar(50)"`
	To       string   `gorm:"type:varchar(50)"`
	CC       string   `gorm:"type:varchar(50)"`
	BCC      string   `gorm:"type:varchar(50)"`
	CronID   int      `gorm:"type:int"`
	Reports  []Report `gorm:"many2many:reports_schedules;"`
}
