package entity

import uuid "github.com/satori/go.uuid"

type LoginResponse struct {
	User  UserResponse  `json:"user"`
	Token TokenResponse `json:"token"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	ID       string     `json:"id"`
	UserName string     `json:"user_name,omitempty"`
	Email    string     `json:"email,omitempty"`
	RoleID   *uuid.UUID `json:"role_id,omitempty"`
}
