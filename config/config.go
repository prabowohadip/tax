package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"sync"
)

var Config config

var onceConfig sync.Once

func init() {
	onceConfig.Do(func() {
		if _, err := toml.DecodeFile("./config.toml", &Config); err != nil {
			fmt.Println(err)
		}
	})
}

type config struct {
	App      appConfig		`toml:"app"`
	Database databaseConfig `toml:"database"`
}

type databaseConfig struct {
	Host   	 string
	Port     string
	Database string
	User     string
	Password string
}

type appConfig struct {
	Name   	 string
	Port     string
	BaseUrl string
}
