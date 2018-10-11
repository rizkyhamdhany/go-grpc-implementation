package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	ENV     string
	Version string
	DB      database `toml:"database"`
}

type database struct {
	DBDriver   string
	DBHost     string
	DBUsername string
	DBPassword string
	DBName     string
	DBCharset  string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}