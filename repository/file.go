package repository

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
	"trail_backend/api/dto"
	"trail_backend/infrastructure"
	"trail_backend/models"
)

type FileRepository interface {
	Create(ctx context.Context, file *models.File) (err error)
	Update(ctx context.Context, file *models.File) (err error)
	GetOneById(ctx context.Context, id string) (res *models.File, err error)
	DeleteById(ctx context.Context, req dto.DeleteFileRequest) (err error)
}

type fileRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewFileRepository(db *infrastructure.Database, logger *zap.Logger) FileRepository {
	return &fileRepository{
		db:     db,
		logger: logger,
	}
}

func (r *fileRepository) Create(ctx context.Context, file *models.File) (err error) {
	err = r.db.Create(&file).Error
	return errors.Wrap(err, "create file failed")
}

func (r *fileRepository) Update(ctx context.Context, file *models.File) (err error) {
	err = r.db.Updates(&file).Error
	return errors.Wrap(err, "update file failed")
}

func (r *fileRepository) GetOneById(ctx context.Context, id string) (res *models.File, err error) {
	var post models.File
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&post).Error; err != nil {
		return nil, errors.Wrap(err, "Find post failed")
	}

	return &post, nil
}

func (r *fileRepository) DeleteById(ctx context.Context, req dto.DeleteFileRequest) (err error) {
	err = r.db.Where("id = ?", req.Id).Updates(map[string]interface{}{"deleted_at": time.Time{}, "updater_id": req.UserId}).Error
	return errors.Wrap(err, "delete file failed")
}
