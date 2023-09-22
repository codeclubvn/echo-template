package dto

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
	UserId string `json:"user_id"`
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

type ListUserResponse struct {
	Data []*UserResponse        `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type GetListUserRequest struct {
	UserId string `json:"user_id"`
	PageOptions
}

type DeleteUserRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

type GetOneUserRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}
