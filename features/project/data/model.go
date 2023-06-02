package data

import (
	"time"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/project"
	task "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/task/data"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ProjectID   string `gorm:"type:varchar(100);primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"size:50;not null;unique"`
	Description string
	StartDate   time.Time
	EndDate     time.Time
	UserID      string
	Tasks       []task.Task `gorm:"foreignKey:ProjectID;references:ProjectID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Project-core to project-model
func projectEntities(p project.Core) Project {
	Project_ID, _ := uuid.NewUUID()
	return Project{
		ProjectID:   Project_ID.String(),
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Name:        p.Name,
		Description: p.Description,
		StartDate:   p.StartDate,
		EndDate:     p.EndDate,
		UserID:      p.UserID,
	}
}
