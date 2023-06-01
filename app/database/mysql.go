package database

import (
	"fmt"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
