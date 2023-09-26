package request

import (
	uuid "github.com/satori/go.uuid"
)

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
	Image   string `json:"image"`
}

type UpdatePostRequest struct {
	ID string `json:"id"`
	CreatePostRequest
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
