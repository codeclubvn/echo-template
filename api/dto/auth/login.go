package dto

import uuid "github.com/satori/go.uuid"

type LoginRequest struct {
	Email       string `json:"email" binding:"required" validate:"email"`
	Password    string `json:"password" binding:"required" validate:"min=6,max=20"`
	RequestFrom string `json:"request_from" binding:"required" enums:"erp/,web,app"`
}

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

type LoginByGoogleRequest struct {
	Email    string `json:"email" binding:"required" validate:"email"`
	GoogleId string `json:"google_id" binding:"required"`
}
