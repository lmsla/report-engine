package entities

type History struct {
	Common
	ID           int    `gorm:"primaryKey;index" json:"id" form:"id"`
	ScheduleID   int    `gorm:"type:int"`
	ScheduleName string `gorm:"type:varchar(50)"`
	// Schedule     Schedule `gorm:"foreignKey:ScheduleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Schedule Schedule `gorm:"foreignKey:ScheduleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// ReportID     int      `gorm:"type:int"`
	// Report       Report   `gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	To          To  `gorm:"serializer:json"`
	CC          CC  `gorm:"serializer:json"`
	BCC         BCC `gorm:"serializer:json"`
	ExecuteTime int64
	EmailTime   int64
	Success     string `gorm:"type:varchar(50)"`
}


