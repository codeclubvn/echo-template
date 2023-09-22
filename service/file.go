package service

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"trail_backend/api/dto"
	"trail_backend/config"
	"trail_backend/models"
	"trail_backend/repository"
	"trail_backend/utils"
)

type (
	FileService interface {
		Create(ctx context.Context, req dto.UploadFileRequest) (*models.File, error)
		Update(ctx context.Context, req dto.UpdateFileRequest) (*models.File, error)
		Delete(ctx context.Context, req dto.DeleteFileRequest) error
		GetOne(ctx context.Context, req dto.GetOneFileRequest) (*models.File, error)
		GetList(ctx context.Context, req dto.GetListFileRequest) (*dto.ListFileResponse, error)
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

func (s *fileService) Create(ctx context.Context, req dto.UploadFileRequest) (*models.File, error) {
	file := &models.File{}
	var err error

	if err = utils.Copy(file, req); err != nil {
		return nil, errors.WithStack(err)
	}
	file.UpdaterID = uuid.FromStringOrNil(req.UserId)

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

	file, err = s.fileRepository.GetOneById(ctx, req)
	return file, err
}

func (s *fileService) GetList(ctx context.Context, req dto.GetListFileRequest) (*dto.ListFileResponse, error) {
	res := &dto.ListFileResponse{}
	var err error

	res, err = s.fileRepository.GetList(ctx, req)
	return res, err
}
