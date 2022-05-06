package config

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	ConfigFile string
	Server     struct {
		Port         string
		SwagerEnable bool
		Timeout      uint32
		Appurl       string
		Appname      string
		LogLevel     uint32 `default:"4"`
	}
	Store struct {
		Host     string `required:"true"`
		Port     int    `required:"true"`
		User     string `required:"true"`
		Password string `required:"true"`
		Dbname   string `required:"true"`
	}
}

func NewConfig(configFile string) (*Config, error) {
	config := &Config{ConfigFile: configFile}
	if err := configor.Load(config, configFile); err != nil {
		return nil, err
	}
	return config, nil
}
