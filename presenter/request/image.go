package request

import (
	"mime/multipart"
)

type UploadImageRequest struct {
	File *multipart.FileHeader `json:"file" swaggerignore:"true" validation:"required"`
}
