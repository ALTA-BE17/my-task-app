package database

import (
	"log"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	book "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book/data"
	user "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user/data"
)

// Add table suffix when creating tables for ACID mysql support by using Engine InnoDB
func Migration(c *config.AppConfig) error {
	db := InitDatabase(c)
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user.User{}, &book.Book{})
	if err != nil {
		log.Fatal(err)
	}

	return err
}
