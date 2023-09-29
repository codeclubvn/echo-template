package utils

import (
	"echo_template/pkg/api_errors"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
	"strings"
)

func CheckListUUID(list pq.StringArray) error {
	for _, v := range list {
		if _, err := uuid.FromString(v); err != nil {
			return errors.New(api_errors.ListFilesInvalid)
		}
	}
	return nil
}

func CheckFileIsImage(file *multipart.FileHeader) error {
	// only accept image: png, jpg, jpeg, gif, svg
	fileNameSplit := strings.Split(file.Filename, ".")
	if len(fileNameSplit) < 2 {
		return errors.New(api_errors.FileIsNotImage)
	}
	extension := fileNameSplit[1]
	if extension != "png" && extension != "jpg" && extension != "jpeg" && extension != "gif" && extension != "svg" {
		return errors.New(api_errors.FileIsNotImage)
	}
	return nil
}
