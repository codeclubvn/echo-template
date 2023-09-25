package service

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"trail_backend/config"
	"trail_backend/dto"
	"trail_backend/models"
	"trail_backend/repository"
	"trail_backend/utils"
)

type (
	FileService interface {
		Upload(ctx context.Context, req dto.UploadFileRequest) (*models.File, error)
		Update(ctx context.Context, req dto.UpdateFileRequest) (*models.File, error)
		Delete(ctx context.Context, req dto.DeleteFileRequest) error
		GetOne(ctx context.Context, req dto.GetOneFileRequest) (*models.File, error)
		Download(ctx context.Context, req dto.DownloadFileRequest) (*models.File, error)
	}
	fileService struct {
		fileRepository repository.FileRepository
		config         *config.Config
	}
)

func NewFileService(itemRepo repository.FileRepository, config *config.Config) FileService {
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

func (s *fileService) Upload(ctx context.Context, req dto.UploadFileRequest) (*models.File, error) {
	file := &models.File{}
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
	file.UpdaterID = uuid.FromStringOrNil(req.UserId)
	file.Path = uploadPath
	file.Size = req.File.Size
	file.ExtensionName = extensionName

	if err = s.fileRepository.Create(ctx, file); err != nil {
		return nil, err
	}

	return file, err
}

func (s *fileService) Update(ctx context.Context, req dto.UpdateFileRequest) (*models.File, error) {
	file := &models.File{}
	var err error

	if err = utils.Copy(file, req); err != nil {
		return nil, err
	}
	file.UpdaterID = uuid.FromStringOrNil(req.UserId)

	if err = s.fileRepository.Update(ctx, file); err != nil {
		return nil, err
	}

	return file, err
}

func (s *fileService) Delete(ctx context.Context, req dto.DeleteFileRequest) error {
	err := s.fileRepository.DeleteById(ctx, req)
	return err
}

func (s *fileService) GetOne(ctx context.Context, req dto.GetOneFileRequest) (*models.File, error) {
	file := &models.File{}
	var err error

	file, err = s.fileRepository.GetOneById(ctx, req.Id)
	return file, err
}

func (s *fileService) Download(ctx context.Context, req dto.DownloadFileRequest) (*models.File, error) {
	file := &models.File{}
	var err error

	file, err = s.fileRepository.GetOneById(ctx, req.Id)
	return file, err
}
