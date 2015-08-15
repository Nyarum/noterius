package core

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
)

// Config struct for config file
type Config struct {
	Base struct {
		IP    string `yaml:"ip"`
		Port  string `yaml:"port"`
		Debug bool   `yaml:"debug"`
		Test  bool   `yaml:"test"`
	} `yaml:"base"`

	Option struct {
		LenBuffer int `yaml:"lenBuffer"`
	} `yaml:"option"`

	Database struct {
		Path string `yaml:"path"`
	} `yaml:"database"`
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
