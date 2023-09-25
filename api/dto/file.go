package dto

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
)

type UploadFileRequest struct {
	UserId   string                `json:"user_id"`
	File     *multipart.FileHeader `json:"file"`
	FileName string                `json:"file_name"`
	Size     int64                 `json:"size"`
	Path     string                `json:"path"`
	Type     string                `json:"type"`
	Data     json.RawMessage       `json:"data"`
}

type UploadFileResponse struct {
	URL string `json:"url"`
}

type UpdateFileRequest struct {
	ID string `json:"id"`
	UploadFileRequest
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

type DownloadFileRequest struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
}
