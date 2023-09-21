package dto

import (
	uuid "github.com/satori/go.uuid"
)

type CreateFileRequest struct {
	UserId      string  `json:"user_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Status      bool    `json:"status"`
	NumberFile  int     `json:"number_product"`
}

type UpdateFileRequest struct {
	ID string `json:"id"`
	CreateFileRequest
}

type FileResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Status      bool      `json:"status"`
	NumberFile  int       `json:"number_product"`
}

type ListFileResponse struct {
	Data []*FileResponse        `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}

type GetListFileRequest struct {
	UserId string `json:"user_id"`
	PageOptions
}

type DeleteFileRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}

type GetOneFileRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}
