package config

type Config struct {
	Port           string `yaml:"port"`
	Server_Host    string `yaml:"server_host"`
	Mongo_Host     string `yaml:"mongo_host"`
	Mongo_Port     string `yaml:"mongo_port"`
	Mongo_User     string `yaml:"mongo_user"`
	Mongo_Password string `yaml:"mongo_password"`
	Mongo_Database string `yaml:"mongo_database"	`
}
