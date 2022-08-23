package entities

type Report struct {
	Common1
	Name        string      `json:"name" form:"name"`
	Description string      `json:"description" form:"description"`
	Datasource  string      `json:"datasource" form:"datasource"`
	Dashboard   string      `json:"dashboard" form:"dashboard"`
	Type        string      `json:"type" form:"type"`
}
