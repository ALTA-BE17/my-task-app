package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	Port      int
	JWTSecret string
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
}

func InitConfig() (*AppConfig, error) {
	return readConfig()
}

func readConfig() (*AppConfig, error) {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("port"); found {
		cnv, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		Port = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("dbusername"); found {
		app.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("dbpassword"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("dbhost"); found {
		app.DB_HOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("dbport"); found {
		cnv, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		app.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("dbname"); found {
		app.DB_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("jwtsecret"); found {
		JWTSecret = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("error while reading config", err.Error())
			return nil, err
		}

		Port = viper.GetInt("port")
		app.DB_USERNAME = viper.GetString("dbusername")
		app.DB_PASSWORD = viper.GetString("dbpassword")
		app.DB_HOST = viper.GetString("dbhost")
		app.DB_PORT = viper.GetInt("dbport")
		app.DB_NAME = viper.GetString("dbname")
		JWTSecret = viper.GetString("jwtsecret")
	}
	return &app, nil
}
