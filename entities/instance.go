package entities

type Instance struct {
	Common
	ID       int    `gorm:"primaryKey;index"`
	Type     string `gorm:"type:varchar(50)"`
	Name     string `gorm:"type:varchar(50)"`
	URL       string `gorm:"type:varchar(50)"`
	User     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50)"`
	Auth     bool   `type:"bool;default:false"`
}
