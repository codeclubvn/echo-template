package repository

import (
	"context"
	"time"
	"trail_backend/api/dto"
	"trail_backend/infrastructure"
	"trail_backend/models"
	"trail_backend/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type FileRepository interface {
	Create(ctx context.Context, product *models.File) (err error)
	Update(ctx context.Context, product *models.File) (err error)
	GetList(ctx context.Context, req dto.GetListFileRequest) (res *dto.ListFileResponse, err error)
	GetOneById(ctx context.Context, req dto.GetOneFileRequest) (res *models.File, err error)
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

func (r *fileRepository) Create(ctx context.Context, product *models.File) (err error) {
	err = r.db.Create(&product).Error
	return errors.Wrap(err, "create product failed")
}

func (r *fileRepository) Update(ctx context.Context, product *models.File) (err error) {
	err = r.db.Updates(&product).Error
	return errors.Wrap(err, "update product failed")
}

func (r *fileRepository) GetOneById(ctx context.Context, req dto.GetOneFileRequest) (res *models.File, err error) {
	var post models.File
	if err := r.db.WithContext(ctx).Where("id = ?", req.Id).First(&post).Error; err != nil {
		return nil, errors.Wrap(err, "Find post failed")
	}

	return &post, nil
}

func (r *fileRepository) GetList(ctx context.Context, req dto.GetListFileRequest) (res *dto.ListFileResponse, err error) {
	var total int64 = 0

	query := r.db.Model(&models.File{})
	if req.Search != "" {
		query = query.Where("name like ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(r.db, req.PageOptions, &res.Data).Count(&total).Error(); err != nil {
		return nil, errors.WithStack(err)
	}

	return res, err
}

func (r *fileRepository) DeleteById(ctx context.Context, req dto.DeleteFileRequest) (err error) {
	err = r.db.Where("id = ?", req.Id).Updates(map[string]interface{}{"deleted_at": time.Time{}, "updater_id": req.UserId}).Error
	return errors.Wrap(err, "delete product failed")
}
