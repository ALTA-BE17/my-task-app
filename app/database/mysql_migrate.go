package database

import (
	"log"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/config"
	project "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/project/data"
	task "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/task/data"
	user "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user/data"
)

// Add table suffix when creating tables for ACID mysql support by using Engine InnoDB
func Migration(c *config.AppConfig) error {
	db := InitDatabase(c)
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&user.User{},
		&project.Project{},
		&task.Task{},
	)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
