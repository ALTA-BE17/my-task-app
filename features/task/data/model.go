package data

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	TaskID      string `gorm:"type:varchar(100);primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"size:50;not null;unique"`
	Description string
	Status      string `gorm:"not null;type:enum('Completed', 'Not Completed'); default:'Not Completed'"`
	ProjectID   string
}
