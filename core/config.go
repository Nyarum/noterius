package core

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"io/ioutil"
	"os"
	pathUtil "path"
)

type ConfigBase struct {
	IP    string `yaml:"ip"`
	Port  string `yaml:"port"`
	Debug bool   `yaml:"test"`
}

type ConfigDatabase struct {
	IP          string `yaml:"ip"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Name        string `yaml:"name"`
	TimeoutSave int    `yaml:"timeout_save"`
}

// Config struct for config file
type Config struct {
	Base     ConfigBase     `yaml:"base"`
	Database ConfigDatabase `yaml:"database"`
}

// NewConfig method for load config
func NewConfig(path string) (config Config, err error) {
	dir, _ := pathUtil.Split(path)
	_, err = os.Stat(path)

	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)

		configNew, err := os.Create(path)
		if err != nil {
			return config, err
		}

		configStat, err := configNew.Stat()
		if err != nil {
			return config, err
		}

		if configStat.Size() == 0 {
			defaultConfig := &Config{
				Base: ConfigBase{
					IP:    "0.0.0.0",
					Port:  "1973",
					Debug: false,
				},
				Database: ConfigDatabase{
					IP:          "127.0.0.1",
					User:        "nota",
					Password:    "notadefault",
					Name:        "noterius",
					TimeoutSave: 15,
				},
			}

			marshalConfig, err := yaml.Marshal(defaultConfig)
			if err != nil {
				return config, err
			}

			_, err = configNew.Write(marshalConfig)
			if err != nil {
				return config, err
			}
		}

		log.WithField("path", path).Info("Default config has been created")
	}

	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(configFile, &config)

	return
}
