package repo

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"trail_backend/domain/repo/model"
	"trail_backend/infra"
)

type FileRepository interface {
	Create(ctx context.Context, file *model.File) (err error)
	Update(ctx context.Context, file *model.File) (err error)
	GetOneById(ctx context.Context, id string) (res *model.File, err error)
	DeleteById(ctx context.Context, id string) (err error)
}

type fileRepository struct {
	db     *infra.Database
	logger *zap.Logger
}

func NewFileRepository(db *infra.Database, logger *zap.Logger) FileRepository {
	return &fileRepository{
		db:     db,
		logger: logger,
	}
}

func (r *fileRepository) Create(ctx context.Context, file *model.File) (err error) {
	err = r.db.Create(&file).Error
	return errors.Wrap(err, "create file failed")
}

func (r *fileRepository) Update(ctx context.Context, file *model.File) (err error) {
	err = r.db.Updates(&file).Error
	return errors.Wrap(err, "update file failed")
}

func (r *fileRepository) GetOneById(ctx context.Context, id string) (res *model.File, err error) {
	var file model.File
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&file).Error; err != nil {
		return nil, errors.Wrap(err, "Find file failed")
	}
	return &file, nil
}

func (r *fileRepository) DeleteById(ctx context.Context, id string) (err error) {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return errors.Wrap(err, "Delete store failed")
	}
	return nil
}
