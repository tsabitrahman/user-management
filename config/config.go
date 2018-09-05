package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//Config hold config data from config.yaml file
type Config struct {
	Port          string `yaml:"port"`
	ServerHost    string `yaml:"server_host"`
	MongoHost     string `yaml:"mongo_host"`
	MongoPort     string `yaml:"mongo_port"`
	MongoUser     string `yaml:"mongo_user"`
	MongoPassword string `yaml:"mongo_password"`
	MongoDatabase string `yaml:"mongo_database"`
}

//GetConfig read configuration file and unmarshal to Config struct
func GetConfig(filepath string) (*Config, error) {
	var config *Config

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
