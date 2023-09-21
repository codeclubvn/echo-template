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

type PostRepository interface {
	Create(ctx context.Context, product *models.Post) (err error)
	Update(ctx context.Context, product *models.Post) (err error)
	GetList(ctx context.Context, req dto.GetListPostRequest) (res *dto.ListPostResponse, err error)
	GetOneById(ctx context.Context, req dto.GetOnePostRequest) (res *models.Post, err error)
	DeleteById(ctx context.Context, req dto.DeletePostRequest) (err error)
}

type postRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewPostRepository(db *infrastructure.Database, logger *zap.Logger) PostRepository {
	return &postRepository{
		db:     db,
		logger: logger,
	}
}

func (r *postRepository) Create(ctx context.Context, product *models.Post) (err error) {
	err = r.db.Create(&product).Error
	return errors.Wrap(err, "create product failed")
}

func (r *postRepository) Update(ctx context.Context, product *models.Post) (err error) {
	err = r.db.Updates(&product).Error
	return errors.Wrap(err, "update product failed")
}

func (r *postRepository) GetOneById(ctx context.Context, req dto.GetOnePostRequest) (res *models.Post, err error) {
	var post models.Post
	if err := r.db.WithContext(ctx).Where("id = ?", req.Id).First(&post).Error; err != nil {
		return nil, errors.Wrap(err, "Find post failed")
	}

	return &post, nil
}

func (r *postRepository) GetList(ctx context.Context, req dto.GetListPostRequest) (res *dto.ListPostResponse, err error) {
	var total int64 = 0

	query := r.db.Model(&models.Post{})
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
func (r *postRepository) DeleteById(ctx context.Context, req dto.DeletePostRequest) (err error) {
	err = r.db.Where("id = ?", req.Id).Updates(map[string]interface{}{"deleted_at": time.Time{}, "updater_id": req.UserId}).Error
	return errors.Wrap(err, "delete product failed")
}
