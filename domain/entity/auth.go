package entity

import uuid "github.com/satori/go.uuid"

type LoginResponse struct {
	User  UserResponse  `json:"user"`
	Token TokenResponse `json:"token"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type UserResponse struct {
	ID       string     `json:"id"`
	UserName string     `json:"user_name"`
	Email    string     `json:"email"`
	RoleID   *uuid.UUID `json:"role_id,omitempty"`
}
