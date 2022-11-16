package entities

type History struct {
	Common
	ScheduleID   int      `gorm:"type:int"`
	ScheduleName string   `gorm:"type:varchar(50)"`
	Schedule     Schedule `gorm:"foreignKey:ScheduleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ReportID     int      `gorm:"type:int"`
	Report       Report   `gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	To           string   `gorm:"type:varchar(50)"`
	CC           string   `gorm:"type:varchar(50)"`
	BCC          string   `gorm:"type:varchar(50)"`
	ExecuteTime  uint
	EmailTime    uint
	Success      bool `gorm:"type:bool;default:false"`
}
