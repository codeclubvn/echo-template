package usecase

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"trail_backend/config"
	"trail_backend/domain/repo"
	"trail_backend/domain/repo/model"
	"trail_backend/pkg/utils"
	"trail_backend/presenter/request"
)

type (
	FileService interface {
		Upload(ctx context.Context, req request.UploadFileRequest) (*model.File, error)
		Update(ctx context.Context, req request.UpdateFileRequest) (*model.File, error)
		Delete(ctx context.Context, id string) error
		GetOne(ctx context.Context, id string) (*model.File, error)
		Download(ctx context.Context, id string) (*model.File, error)
	}
	fileService struct {
		fileRepository repo.FileRepository
		config         *config.Config
	}
)

func NewFileService(itemRepo repo.FileRepository, config *config.Config) FileService {
	return &fileService{
		fileRepository: itemRepo,
		config:         config,
	}
}

func createFolder(fileId string, config *config.Config) string {
	firstChar := fileId[0:1]
	secondChar := fileId[1:2]
	uploadPath := config.Server.UploadPath + "/" + firstChar + "/" + secondChar + "/"

	// create folder if not exists
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadPath, 0755); err != nil {
			panic(err)
		}
	}
	return uploadPath
}

func saveToFolder(file *multipart.FileHeader, uploadPath, id, extensionName string) error {
	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(uploadPath + id + extensionName)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func getExtensionNameFromFilename(fileName string) string {
	return strings.Split(fileName, ".")[1]
}

func (s *fileService) Upload(ctx context.Context, req request.UploadFileRequest) (*model.File, error) {
	file := &model.File{}
	var err error
	fileId := uuid.NewV4().String()
	extensionName := getExtensionNameFromFilename(req.FileName)

	uploadPath := createFolder(fileId, s.config)
	if err = saveToFolder(req.File, uploadPath, fileId, extensionName); err != nil {
		return nil, errors.WithStack(err)
	}

	if err = utils.Copy(file, req); err != nil {
		return nil, errors.WithStack(err)
	}
	file.ID = uuid.FromStringOrNil(fileId)
	file.Path = uploadPath
	file.Size = req.File.Size
	file.ExtensionName = extensionName

	if err = s.fileRepository.Create(ctx, file); err != nil {
		return nil, err
	}

	return file, err
}

func (s *fileService) Update(ctx context.Context, req request.UpdateFileRequest) (*model.File, error) {
	file := &model.File{}
	var err error

	if err = utils.Copy(file, req); err != nil {
		return nil, err
	}
	if err = s.fileRepository.Update(ctx, file); err != nil {
		return nil, err
	}

	return file, err
}

func (s *fileService) Delete(ctx context.Context, id string) error {
	return s.fileRepository.DeleteById(ctx, id)
}

func (s *fileService) GetOne(ctx context.Context, id string) (*model.File, error) {
	return s.fileRepository.GetOneById(ctx, id)
}

func (s *fileService) Download(ctx context.Context, id string) (*model.File, error) {
	return s.fileRepository.GetOneById(ctx, id)
}
