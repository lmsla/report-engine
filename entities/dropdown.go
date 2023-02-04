package entities

type Dropdown struct {
	Common
	text  string `gorm:"type:varchar(50)"`
	value string `gorm:"type:varchar(50)"`
}
