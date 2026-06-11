package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName string `envconfig:"APP_NAME"`
	AppPort string `envconfig:"APP_PORT"`

	DbName         string `envconfig:"DB_NAME"`
	DBSchema       string `envconfig:"DB_SCHEMA"`
	DbUser         string `envconfig:"DB_USERNAME"`
	DbPassword     string `envconfig:"DB_PASSWORD"`
	DbHost         string `envconfig:"DB_HOST"`
	DbPort         string `envconfig:"DB_PORT"`
	DbMaxIdleConns int    `envconfig:"db_max_idle_conns" default:"50"`
	DbMaxOpenConns int    `envconfig:"db_max_open_conns" default:"100"`
}

var once sync.Once
var instance Config

func GetConfig() Config {
	once.Do(func() {
		err := envconfig.Process("", &instance)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return instance
}

func SetConfig(cfg Config) {
	instance = cfg
}
