package models

type Report struct {
	Common
	ID         int       `json:"id" form:"id"`
	Name       string    `json:"name" form:"name"`
	TimeUnit   int       `json:"time_unit" form:"time_unit"`
	TimePeriod string    `json:"time_period" form:"time_period"`
	Elements   []Element `json:"elements" form:"elements"`
	// Elements   []Element `gorm:"foreignKey:ReportID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`

}

type Element struct {
	Common
	ID         int    `json:"id" form:"id"`
	ReportID   int    `json:"report_id" form:"report_id"`
	Type       string `json:"type" form:"type"`
	Name       string `json:"name" form:"name"`
	UID        string `json:"uid" form:"uid"`
	RowNum     int    `json:"row_num" form:"row_num"`
	ColumnType string `json:"column_type" form:"column_type"`
	InstanceID int    `json:"instance_id" form:"instance_id"`
	// Instance   Instance `json:"instance" form:"instance"`
	SpcaceName string `json:"space_name" form:"space_name"`
}
