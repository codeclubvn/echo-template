package dto

import (
	uuid "github.com/satori/go.uuid"
)

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
	Image   string `json:"image"`
	UserId  string `json:"user_id"`
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
	UserId string `json:"user_id"`
	PageOptions
}

type ListPostResponse struct {
	Data []*PostResponse        `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type DeletePostRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

type GetOnePostRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}
