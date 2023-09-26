package usecase

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"trail_backend/config"
	"trail_backend/domain/repo"
	"trail_backend/presenter/request"
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
