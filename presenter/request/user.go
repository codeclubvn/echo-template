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

type DeleteUserRequest struct {
	ID     string    `json:"id"`
	UserId uuid.UUID `json:"user_id" swaggerignore:"true"`
}

type GetListUserRequest struct {
	PageOptions
}
