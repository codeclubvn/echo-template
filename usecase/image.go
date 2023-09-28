package usecase

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"trial_backend/config"
	"trial_backend/domain/repo"
	"trial_backend/presenter/request"
)

type (
	FileCloudService interface {
		Update(ctx context.Context, req request.UploadFileRequest) (*uploader.UploadResult, error)
	}
	imageService struct {
		imageRepository repo.CloudinaryRepository
		config          *config.Config
	}
)

func NewImageService(itemRepo repo.CloudinaryRepository, config *config.Config) FileCloudService {
	return &imageService{
		imageRepository: itemRepo,
		config:          config,
	}
}

func (s *imageService) Update(ctx context.Context, req request.UploadFileRequest) (*uploader.UploadResult, error) {
	res, err := s.imageRepository.UploadFileCloud(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, err
}
