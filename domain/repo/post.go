package repo

import (
	"context"
	"echo_template/domain/repo/model"
	"echo_template/infra"
	"echo_template/pkg/api_errors"
	"echo_template/pkg/constants"
	"echo_template/pkg/utils"
	"echo_template/presenter/request"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) error
	Update(ctx context.Context, post *model.Post) error
	GetList(ctx context.Context, req request.GetListPostRequest) ([]*model.Post, *int64, error)
	GetOneById(ctx context.Context, id string) (*model.Post, error)
	DeleteById(ctx context.Context, id string) error
}

type postRepository struct {
	db     *infra.Database
	logger *zap.Logger
}

func NewPostRepository(db *infra.Database, logger *zap.Logger) PostRepository {
	return &postRepository{
		db:     db,
		logger: logger,
	}
}

func (r *postRepository) Create(ctx context.Context, post *model.Post) error {
	err := r.db.Create(&post).Error
	return errors.Wrap(err, "create post failed")
}

func (r *postRepository) Update(ctx context.Context, post *model.Post) error {
	err := r.db.Updates(&post).Error
	return errors.Wrap(err, "update post failed")
}

func (r *postRepository) GetOneById(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&post).Error; err != nil {
		return nil, errors.New(api_errors.ErrPostNotFound)
	}

	return &post, nil
}

func (r *postRepository) GetList(ctx context.Context, req request.GetListPostRequest) ([]*model.Post, *int64, error) {
	var res []*model.Post
	var total int64 = 0

	query := r.db.Model(&model.Post{})
	if req.Search != "" {
		search := "%" + req.Search + "%"
		query = query.Where(`title ilike ? or SIMILARITY(unaccent(title),?) > 0.25`, search, req.Search)
	}

	switch req.Sort {
	case constants.SortCreatedAsc:
		query = query.Order(constants.SortCreatedAsc)
	default:
		query = query.Order(constants.SortCreatedDesc)
	}

	if err := utils.QueryPagination(query, req.PageOptions, &res); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return res, &total, nil
}

func (r *postRepository) DeleteById(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Post{}).Error; err != nil {
		return errors.Wrap(err, "Delete post failed")
	}
	return nil
}
