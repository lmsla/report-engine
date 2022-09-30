package entities

type Schedule struct {
	Common
	ScheduleID     int    `json:"schedule_id" form:"schedule_id" gorm:"primaryKey"`
	Name           string `json:"schedule_name" form:"schedule_name"`
	Description    string `json:"description" form:"description"`
	// Report         string `json:"reports" form:"reports"`
	GenerateReport string `json:"generate-report" form:"generate_report"`
	SendReport     string `json:"send_report" form:"send_report"`
	Recipient      string `json:"recipients" form:"recipients"`
	Status         string `json:"status" form:"status"`
	Report []Report `gorm:"foreignKey:ScheduleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

}

type Status struct {
	Status string `json:"status" form:"status"`
}

type FileHistory struct {
	Common
	FileID     int    `json:"file_id" form:"file_id" gorm:"primaryKey"`
	ScheduleID int    `json:"schedule_id" form:"schedule_id"`
	HtmlName   string `json:"html_name" form:"html_name"`
	PdfName    string `json:"pdf_name" form:"pdf_name"`
	Filetype   string `json:"file_type" form:"file_type"`
}
