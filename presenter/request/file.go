package request

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
)

type UploadFileRequest struct {
	File     *multipart.FileHeader `json:"file" swaggerignore:"true" validation:"required"`
	FileName string                `json:"file_name"`
	Size     int64                 `json:"size"`
	Path     string                `json:"path"`
	Type     string                `json:"type"`
	Data     json.RawMessage       `json:"data,omitempty" swaggertype:"array,string"`
}

type UploadFileResponse struct {
	URL string `json:"url"`
}

type UpdateFileRequest struct {
	ID       string                `json:"id" form:"id" validate:"required"`
	File     *multipart.FileHeader `json:"file" swaggerignore:"true"`
	FileName string                `json:"file_name" form:"file_name"`
	Data     json.RawMessage       `json:"data,omitempty" swaggertype:"array,string"`
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
