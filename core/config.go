package core

import "github.com/BurntSushi/toml"

type CommonSub struct {
	Host string `toml:"host"`
}

type DatabaseSub struct {
	Dsn string `toml:"dsn"`
}

type Config struct {
	Common   CommonSub   `toml:"common"`
	Database DatabaseSub `toml:"database"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load(path string) error {
	_, err := toml.DecodeFile(path, c)

	return err
}
