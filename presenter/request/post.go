package request

import (
	uuid "github.com/satori/go.uuid"
)

type CreatePostRequest struct {
	UserId  uuid.UUID   `json:"user_id" swaggerignore:"true"`
	Title   string      `json:"title"`
	Content string      `json:"content"`
	Slug    string      `json:"slug"`
	Image   string      `json:"image"`
	Files   FileIdSlice `json:"files" gorm:"column:files;type:uuid[]"`
}

type FileIdSlice []uuid.UUID

type UpdatePostRequest struct {
	ID string `json:"id"`
	CreatePostRequest
}

type DeletePostRequest struct {
	ID     string    `json:"id"`
	UserId uuid.UUID `json:"user_id" swaggerignore:"true"`
}

type PostResponse struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Slug    string    `json:"slug"`
	Image   string    `json:"image"`
}

type GetListPostRequest struct {
	PageOptions
}
