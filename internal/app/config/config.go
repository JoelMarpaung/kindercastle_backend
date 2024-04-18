package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

//go:generate easytags $GOFILE envconfig

type Config struct {
	AppPort string
	AppEnv  string

	DBName    string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBMaxOpen int
	DBMaxIdle int

	EnableDocs       bool
	FirebaseCredPath string
}

var (
	once sync.Once
	conf Config
)

func GetConfig() Config {
	once.Do(func() {
		if os.Getenv("APP_ENV") != "production" {
			if err := godotenv.Load(); err != nil {
				log.Printf("No .env file found or error loading .env file")
			}
		}

		conf.AppPort = os.Getenv("PORT")
		conf.AppEnv = os.Getenv("APP_ENV")
		conf.DBName = os.Getenv("DB_NAME")
		conf.DBHost = os.Getenv("DB_HOST")
		conf.DBPort = os.Getenv("DB_PORT")
		conf.DBUser = os.Getenv("DB_USER")
		conf.DBPass = os.Getenv("DB_PASS")
		conf.DBMaxOpen = getIntEnv("DB_MAX_OPEN", 10)
		conf.DBMaxIdle = getIntEnv("DB_MAX_IDLE", 10)

		conf.EnableDocs = getBoolEnv("ENABLE_DOCS", false)
		conf.FirebaseCredPath = os.Getenv("FIREBASE_CRED_PATH")
	})

	return conf
}

func getIntEnv(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		boolValue, err := strconv.ParseBool(value)
		if err == nil {
			return boolValue
		}
	}
	return defaultValue
}
