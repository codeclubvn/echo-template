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
	Create(ctx context.Context, post *models.Post) (err error)
	Update(ctx context.Context, post *models.Post) (err error)
	GetList(ctx context.Context, req dto.GetListPostRequest) (res []*models.Post, total *int64, err error)
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

func (r *postRepository) Create(ctx context.Context, post *models.Post) (err error) {
	err = r.db.Create(&post).Error
	return errors.Wrap(err, "create post failed")
}

func (r *postRepository) Update(ctx context.Context, post *models.Post) (err error) {
	err = r.db.Updates(&post).Error
	return errors.Wrap(err, "update post failed")
}

func (r *postRepository) GetOneById(ctx context.Context, req dto.GetOnePostRequest) (res *models.Post, err error) {
	var post models.Post
	if err := r.db.WithContext(ctx).Where("id = ?", req.Id).First(&post).Error; err != nil {
		return nil, errors.Wrap(err, "Find post failed")
	}

	return &post, nil
}

func (r *postRepository) GetList(ctx context.Context, req dto.GetListPostRequest) (res []*models.Post, total *int64, err error) {
	query := r.db.Model(&models.Post{})
	if req.Search != "" {
		query = query.Where("name like ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res); err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if err = query.Count(total).Error; err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return res, total, err
}
func (r *postRepository) DeleteById(ctx context.Context, req dto.DeletePostRequest) (err error) {
	err = r.db.Where("id = ?", req.Id).Updates(map[string]interface{}{"deleted_at": time.Time{}, "updater_id": req.UserId}).Error
	return errors.Wrap(err, "delete post failed")
}
