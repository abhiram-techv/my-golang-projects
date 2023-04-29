package config

import (
	"io/ioutil"
	"myproject/utils"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig() Config {
	var config Config
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		utils.Logging("failed to read config file", err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		utils.Logging("failed to unmarshal config file contents", err)
	}
	return config
}
