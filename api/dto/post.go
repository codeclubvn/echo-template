package dto

import (
	uuid "github.com/satori/go.uuid"
)

type CreatePostRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Status      bool    `json:"status"`
	NumberPost  int     `json:"number_product"`
	UserId      string  `json:"user_id"`
}

type UpdatePostRequest struct {
	ID string `json:"id"`
	CreatePostRequest
}

type PostResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Status      bool      `json:"status"`
	NumberPost  int       `json:"number_product"`
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
