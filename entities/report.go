package entities

type Report struct {
	Common
	ID         int        `gorm:"primaryKey;index"`
	Name       string     `gorm:"type:varchar(50)"`
	TimeUnit   string     `gorm:"type:varchar(50)"`
	TimePeriod int        `gorm:"type:int"`
	Elements   []Element  `gorm:"foreignKey:ReportID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Schedules  []Schedule `gorm:"many2many:reports_schedules;"`
}

type Element struct {
	Common
	ID         int      `gorm:"primaryKey;index"`
	ReportID   int      `gorm:"index"`
	Type       string   `gorm:"type:varchar(50)"`
	Name       string   `gorm:"type:varchar(50)"`
	UID        string   `gorm:"type:varchar(50)"`
	RowNum     int      `gorm:"type:int"`
	ColumnType string   `gorm:"type:varchar(50)"`
	InstanceID int      `gorm:"type:int"`
	Instance   Instance `gorm:"foreignKey:InstanceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpcaceName string   `gorm:"type:varchar(50)"`
}
