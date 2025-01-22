package entities

type DropdownBody struct {
	SpaceName  string `json:"space_name" form:"space_name"`
	InstanceID int    `json:"instance_id" form:"instance_id"`
	SourceType string `json:"source_type" form:"source_type"`
}

type FieldsDropdownBody struct {
	SpaceName  string `json:"space_name" form:"space_name"`
	InstanceID int    `json:"instance_id" form:"instance_id"`
	DataViewID string `json:"data_view_id" form:"data_view_id"`
}

type Dropdown struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}