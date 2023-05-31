package database

import (
	"fmt"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.DB_USERNAME, config.Database.DB_PASSWORD, config.Database.DB_HOST, config.Database.DB_PORT, config.Database.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
