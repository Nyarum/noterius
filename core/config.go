package core

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
)

// Config struct for config file
type Config struct {
	IP   string `yaml:"IP"`
	Port string `yaml:"Port"`
}

// LoadConfig function for load config
func LoadConfig(config *Config, path string) (err error) {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(configFile, config)

	return
}
