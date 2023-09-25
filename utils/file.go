package utils

import (
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"trail_backend/dto"
)

func GetFile(c echo.Context, req *dto.UploadFileRequest, folder string) error {
	file, err := c.FormFile("file")
	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(folder + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	req.FileName = folder + file.Filename
	req.Size = file.Size
	return nil
}
