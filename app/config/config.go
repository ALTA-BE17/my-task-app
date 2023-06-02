package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	JWTSECRET string
)

type AppConfig struct {
	DBUSERNAME string
	DBPASSWORD string
	DBHOST     string
	DBPORT     int
	DBNAME     string
}

func InitConfig() (*AppConfig, error) {
	return readConfig()
}

func readConfig() (*AppConfig, error) {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("DBUSERNAME"); found {
		app.DBUSERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASSWORD"); found {
		app.DBPASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		app.DBPORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("JWTSECRET"); found {
		JWTSECRET = val
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

		app.DBUSERNAME = viper.GetString("DBUSERNAME")
		app.DBPASSWORD = viper.GetString("DBPASSWORD")
		app.DBHOST = viper.GetString("DBHOST")
		app.DBPORT = viper.GetInt("DBPORT")
		app.DBNAME = viper.GetString("DBNAME")
		JWTSECRET = viper.GetString("JWTSECRET")
	}
	return &app, nil
}
