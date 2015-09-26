package core

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
)

// Config struct for config file
type Config struct {
	Base struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
		Test bool   `yaml:"test"`
	} `yaml:"base"`

	Database struct {
		IP          string `yaml:"ip"`
		User        string `yaml:"user"`
		Password    string `yaml:"password"`
		Name        string `yaml:"name"`
		TimeoutSave int    `yaml:"timeout_save"`
	} `yaml:"database"`
}

// NewConfig method for load config
func NewConfig(config *Config, path string) (err error) {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(configFile, config)

	return
}
