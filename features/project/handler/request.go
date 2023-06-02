package handler

import (
	"time"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/project"
)

type CreateProjectRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	StartDate   string `json:"start_date" form:"start_date"`
	EndDate     string `json:"end_date" form:"end_date"`
}

func RequestToCore(data interface{}) *project.Core {
	res := project.Core{}
	switch v := data.(type) {
	case CreateProjectRequest:
		res.Name = v.Name
		res.Description = v.Description
		res.StartDate, _ = time.Parse("2006-01-02", v.StartDate)
		res.EndDate, _ = time.Parse("2006-01-02", v.EndDate)
	default:
		return nil
	}
	return &res
}
