package request

import (
	uuid "github.com/satori/go.uuid"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	Social   string `json:"social"`
	SocialId string `json:"social_id"`
}

type UpdateUserRequest struct {
	Id string `json:"id"`
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

type GetListUserRequest struct {
	PageOptions
}
