package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

//go:generate easytags $GOFILE envconfig

type Config struct {
	AppPort string `envconfig:"app_port" default:"1337"`
	AppEnv  string `envconfig:"app_env"`

	DBName    string `envconfig:"db_name"`
	DBHost    string `envconfig:"db_host"`
	DBPort    string `envconfig:"db_port"`
	DBUser    string `envconfig:"db_user"`
	DBPass    string `envconfig:"db_pass"`
	DBMaxOpen int    `envconfig:"db_max_open" default:"10"`
	DBMaxIdle int    `envconfig:"db_max_idle" default:"10"`

	EnableDocs bool `envconfig:"enable_docs" default:"false"`
}

var (
	once sync.Once
	conf Config
)

func GetConfig() Config {
	once.Do(func() {
		_ = godotenv.Load()

		err := envconfig.Process("", &conf)
		if err != nil {
			log.Fatal(err)
		}
	})

	return conf
}
