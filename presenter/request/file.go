package request

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
	Data     json.RawMessage       `json:"domain"`
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
