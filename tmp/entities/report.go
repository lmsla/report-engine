package entities

type Report struct {
	Common
	ID         int        `gorm:"primaryKey;index" json:"id" form:"id"`
	Name       string     `gorm:"type:varchar(50)" json:"name" form:"name"`
	TimeUnit   string     `gorm:"type:varchar(50)" json:"time_unit" form:"time_unit"`
	TimePeriod int        `gorm:"type:int" json:"time_period" form:"time_period"`
	Alias      string     `gorm:"type:varchar(45)" json:"alias" form:"alias"`
	Elements   []Element  `gorm:"foreignKey:ReportID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Schedules  []Schedule `gorm:"many2many:reports_schedules;"`
	// DataTables []DataTable `gorm:"foreignKey:ReportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tables []Table `gorm:"many2many:reports_tables;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;foreignKey:ID;reference:ID;"`
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
	Instance   Instance `json:"instance" gorm:"foreignKey:InstanceID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpaceName  string   `gorm:"type:varchar(50)" json:"space_name" form:"space_name"`
}

type Table struct {
	Common
	ID int `gorm:"primaryKey;index" json:"id" form:"id"`
	// ReportID   int      `gorm:"index" json:"report_id" form:"report_id"`
	Type       string   `gorm:"type:varchar(50)" json:"type" form:"type"`
	Name       string   `gorm:"type:varchar(50)" json:"name" form:"name"`
	DataView   string   `gorm:"type:varchar(50)" json:"data_view" form:"data_view"`
	UID        string   `gorm:"type:varchar(50)" json:"uid" form:"uid"`
	RowNum     int      `gorm:"type:int" json:"row_num" form:"row_num"`
	InstanceID int      `gorm:"type:int" json:"instance_id" form:"instance_id"`
	Instance   Instance `json:"instance" gorm:"foreignKey:InstanceID;reference:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SpaceName  string   `gorm:"type:varchar(50)" json:"space_name" form:"space_name"`
	Reports    []Report `gorm:"many2many:reports_tables;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Columns    []Column `gorm:"foreignKey:TableID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Column struct {
	Common
	ID      int    `gorm:"primaryKey;index" json:"id" form:"id"`
	TableID int    `gorm:"index" json:"table_id" form:"table_id"`
	Name    string `gorm:"type:varchar(50)" json:"name" form:"name"`
	Alias   string `gorm:"type:varchar(50)" json:"alias" form:"alias"`
	Order   int    `gorm:"type:int" json:"order" form:"order"`
	Size    int    `gorm:"type:int" json:"size" form:"size"`
}

type ReportsTables struct {
	Common
	ReportID int `gorm:"primaryKey" form:"report_id"`
	TableID  int `gorm:"primaryKey" form:"table_id"`
}

type ScreenshotUrl struct {
	User     string `json:"user" form:"user"`
	Password string `json:"password" form:"password"`
	Url      string `json:"url" form:"url"`
	Name     string `json:"name" form:"name"`
}
