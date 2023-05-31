package handler

import "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"

type InsertBookRequest struct {
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}

type UpdateBookRequest struct {
	ID     uint
	Title  *string `json:"title" form:"title"`
	Author *string `json:"author" form:"author"`
}

func RequestToCore(data interface{}) *book.Core {
	res := book.Core{}
	switch v := data.(type) {
	case InsertBookRequest:
		res.Title = v.Title
		res.Author = v.Author
	case *UpdateBookRequest:
		res.ID = v.ID
		if v.Title != nil {
			res.Title = *v.Title
		}
		if v.Author != nil {
			res.Author = *v.Author
		}
	default:
		return nil
	}
	return &res
}
