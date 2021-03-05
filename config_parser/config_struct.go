package config_parser

type ConfigStructure struct {
	Contents ConfigCore `json:"configuration"`
}

type ConfigCore struct {
	Database string `json:"database"`
	WebPath string `json:"webPath"`
	IDP ConfigIDP `json:"idp"`
}

type ConfigIDP struct {
	AuthURL string `json:"authUrl"`
	ClientID string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}