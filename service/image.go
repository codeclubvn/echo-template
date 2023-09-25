package service

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"trail_backend/config"
	"trail_backend/dto"
	"trail_backend/repository"
)

type (
	FileCloudService interface {
		Update(ctx context.Context, req dto.UploadFileRequest) (*uploader.UploadResult, error)
	}
	imageService struct {
		imageRepository repository.CloudinaryRepository
		config          *config.Config
	}
)

func NewImageService(itemRepo repository.CloudinaryRepository, config *config.Config) FileCloudService {
	return &imageService{
		imageRepository: itemRepo,
		config:          config,
	}
}

func (s *imageService) Update(ctx context.Context, req dto.UploadFileRequest) (*uploader.UploadResult, error) {
	res, err := s.imageRepository.UploadFileCloud(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, err
}
