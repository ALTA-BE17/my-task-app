package database

import (
	"fmt"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.DBUSERNAME, config.DBPASSWORD, config.DBHOST, config.DBPORT, config.DBNAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
