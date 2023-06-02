package data

import (
	"time"

	task "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/task/data"
	"gorm.io/gorm"
)

type Project struct {
	ProjectID   string `gorm:"type:varchar(100);primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	UserID      string
	Tasks       []task.Task `gorm:"foreignKey:ProjectID;references:ProjectID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
