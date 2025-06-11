package models

type Report struct {
	Common
	ID         int    `gorm:"primaryKey;index" json:"id" form:"id"`
	Name       string `gorm:"type:varchar(50)" json:"name" form:"name"`
	TimeUnit   string `gorm:"type:varchar(50)" json:"time_unit" form:"time_unit"`
	TimePeriod int    `gorm:"type:int" json:"time_period" form:"time_period"`
	Alias      string `gorm:"type:varchar(50)" json:"alias" form:"alias"`
	Elements   []Element `gorm:"foreignKey:ReportID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// Schedules  []Schedule `gorm:"many2many:reports_schedules;"`
}

type Element struct {
	Common
	ID         int      `gorm:"primaryKey;index" json:"id" form:"id"`
	ReportID   int      `gorm:"index" json:"report_id" form:"report_id"`
	Type       string   `gorm:"type:varchar(50)" json:"type" form:"type"`
	Name       string   `gorm:"type:varchar(50)" json:"name" form:"name"`
	UID        string   `gorm:"type:varchar(50)" json:"uid" form:"uid"`
	RowNum     int      `gorm:"type:int" json:"row_num" form:"row_num"`
	ColumnType string   `gorm:"type:varchar(50)" json:"column_type" form:"column_type"`
	InstanceID int      `gorm:"type:int" json:"instance_id" form:"instance_id"`
	Instance   Instance `gorm:"foreignKey:InstanceID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpaceName  string   `gorm:"type:varchar(50)" json:"space_name" form:"space_name"`
}

