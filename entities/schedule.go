package entities

type Schedule struct {
	Common1
	Name           string `json:"schedule_name" form:"schedule_name"`
	Description    string `json:"description" form:"description"`
	Report         string `json:"reports" form:"reports"`
	GenerateReport string `json:"generate-report" form:"generate_report"`
	SendReport     string `json:"send_report" form:"send_report"`
	Recipient      string `json:"recipients" form:"recipients"`
	Status         string `json:"status" form:"status"`
}
