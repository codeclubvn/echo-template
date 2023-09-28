package usecase

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
	"trial_backend/config"
	"trial_backend/domain/repo"
)

type (
	FileCloudService interface {
		Update(ctx context.Context, file *multipart.FileHeader) (*uploader.UploadResult, error)
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

func (s *imageService) Update(ctx context.Context, file *multipart.FileHeader) (*uploader.UploadResult, error) {
	res, err := s.imageRepository.UploadFileCloud(ctx, file)
	if err != nil {
		return nil, err
	}

	return res, err
}
