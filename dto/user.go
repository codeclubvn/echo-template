package dto

import (
	uuid "github.com/satori/go.uuid"
)

type CreateUserRequest struct {
	UserId      string  `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Status      bool    `json:"status"`
	NumberUser  int     `json:"number_product"`
}

type UpdateUserRequest struct {
	ID string `json:"id"`
	CreateUserRequest
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Status      bool      `json:"status"`
	NumberUser  int       `json:"number_product"`
}

type UsersResponse struct {
	Data []*UserResponse        `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type ListUserRequest struct {
	PageOptions
}
