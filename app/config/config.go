package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var JWTSecret string

type Database struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
}

type AppConfig struct {
	Port      int
	Database  Database
	JWTSecret string
}

func InitConfig() *AppConfig {
	return readConfig()
}

func readConfig() *AppConfig {
	// inisialisasi variabel dg type struct AppConfig
	app := AppConfig{}
	isRead := true

	// proses mencari & membaca environment var dg key tertentu
	if val, found := os.LookupEnv("jwtsecret"); found {
		app.JWTSecret = val
		isRead = false
	}
	if val, found := os.LookupEnv("port"); found {
		cnv, _ := strconv.Atoi(val)
		app.Port = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("database.username"); found {
		app.Database.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("database.password"); found {
		app.Database.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("database.host"); found {
		app.Database.DB_HOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("database.port"); found {
		cnv, _ := strconv.Atoi(val)
		app.Database.DB_PORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("database.name"); found {
		app.Database.DB_NAME = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("error while reading config", err.Error())
			return nil
		}

		app := AppConfig{
			Port: viper.GetInt("port"),
			Database: Database{
				DB_USERNAME: viper.GetString("database.username"),
				DB_PASSWORD: viper.GetString("database.password"),
				DB_HOST:     viper.GetString("database.host"),
				DB_PORT:     viper.GetInt("database.port"),
				DB_NAME:     viper.GetString("database.name"),
			},
			JWTSecret: viper.GetString("jwt"),
		}
		return &app
	}
	return &app
}
