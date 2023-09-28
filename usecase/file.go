package usecase

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"trial_backend/config"
	"trial_backend/domain/repo"
	"trial_backend/domain/repo/model"
	"trial_backend/pkg/constants"
	"trial_backend/pkg/utils"
	"trial_backend/presenter/request"
)

type (
	FileService interface {
		Upload(ctx context.Context, file *multipart.FileHeader) (*model.File, error)
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
	dst, err := os.Create(uploadPath + id + "." + extensionName)
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
	fileNameArr := strings.Split(fileName, ".")
	extensionName := ""
	if len(fileNameArr) > constants.NumberFileNameSplit {
		extensionName = fileNameArr[1]
	}
	return extensionName
}

func (s *fileService) Upload(ctx context.Context, req *multipart.FileHeader) (*model.File, error) {
	var err error
	fileId := uuid.NewV4().String()
	extensionName := getExtensionNameFromFilename(req.Filename)

	uploadPath := createFolder(fileId, s.config)
	if err = saveToFolder(req, uploadPath, fileId, extensionName); err != nil {
		return nil, errors.WithStack(err)
	}

	file := &model.File{
		BaseModel: model.BaseModel{
			ID: uuid.FromStringOrNil(fileId),
		},
		Path:          uploadPath,
		Size:          req.Size,
		ExtensionName: extensionName,
		FileName:      req.Filename,
	}
	if err = s.fileRepository.Create(ctx, file); err != nil {
		return nil, err
	}

	return file, err
}

func (s *fileService) Update(ctx context.Context, req request.UpdateFileRequest) (*model.File, error) {
	// get one file
	file, err := s.fileRepository.GetOneById(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// check filePath is not in ./domain/assets
	if !strings.Contains(file.Path, s.config.Server.UploadPath) {
		file.Path = createFolder(file.ID.String(), s.config)
	} else {
		// delete old file
		_ = os.Remove(file.Path + file.ID.String() + "." + file.ExtensionName)
	}

	if err = utils.Copy(file, req); err != nil {
		return nil, err
	}
	file.ExtensionName = getExtensionNameFromFilename(req.File.Filename)
	file.Size = req.File.Size

	// create folder if not exists
	if _, err := os.Stat(file.Path); os.IsNotExist(err) {
		if err := os.MkdirAll(file.Path, 0755); err != nil {
			panic(err)
		}
	}
	if err := saveToFolder(req.File, file.Path, file.ID.String(), file.ExtensionName); err != nil {
		return nil, errors.WithStack(err)
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
