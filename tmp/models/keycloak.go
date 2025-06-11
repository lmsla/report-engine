package models

// import "github.com/Nerzal/gocloak"

type SSOUser struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Realm       string          `json:"realm"`
	Groups      []RealmGroupRes `json:"groups"`
	Roles       []string        `json:"roles"`
	IsAdmin     bool            `json:"is_admin"`
	AccessHosts []string        `json:"access_hosts"`
}

// ******************************      Group     ******************************//
type Realm struct {
	Common
	Name           string          `json:"name"   form:"realm_name" gorm:"primaryKey"` // from sso
	RealmGroups    []RealmGroup    `json:"realm_groups"  form:"realm_groups" gorm:"foreignKey:RealmName;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RealmGroup struct {
	Common
	RealmName       string        `json:""           form:"realm_name"` // from sso
	ID              string        `json:"id"  form:"id"`                // from sso
	ResourceGroupID int           `json:"resource_group_id"   form:"resource_group_id" gorm:"index:,unique"`
}


type RealmGroupRes struct {
	ID     string             `json:"id" `
	Name   string             `json:"name"`
}