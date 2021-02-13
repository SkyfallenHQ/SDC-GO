package config_parser

type config_file struct {

	config []config_json `json:"config"`

}

type config_json struct {

	database config_database `json:"database"`

}

type config_database struct {

	mongouri string `json:"mongouri"`

}