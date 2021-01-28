package types

import "gorm.io/datatypes"

type Configuration struct {
	Domains []Domain
}

type Domain struct {
	SiteID  int      `json:"site_id"`
	Name    string   `json:"name"`
	Plugins []Plugin `json:"plugins"`
}

type Plugin struct {
	ID        int    `json:"id"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Settings  datatypes.JSON
}
