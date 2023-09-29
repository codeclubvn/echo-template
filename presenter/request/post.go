package request

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type CreatePostRequest struct {
	UserId  uuid.UUID      `json:"user_id" swaggerignore:"true"`
	Title   string         `json:"title"`
	Content string         `json:"content"`
	Slug    string         `json:"slug"`
	Image   string         `json:"image"`
	Files   pq.StringArray `json:"files" swaggertype:"array,string" example:"52bdcd0a-5615-430b-bfc4-89fc40bd6b71"`
}

type UpdatePostRequest struct {
	ID string `json:"id" validate:"required"`
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
