package repo

import (
	"context"
	"echo_template/domain/repo/model"
	"echo_template/infra"
	"echo_template/pkg/api_errors"
	"echo_template/pkg/utils"
	"echo_template/presenter/request"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetList(ctx context.Context, req request.GetListUserRequest) ([]*model.User, *int64, error)
	GetOneById(ctx context.Context, id string) (*model.User, error)
	DeleteById(ctx context.Context, id string) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetBySocialId(ctx context.Context, socialId string) (*model.User, error)
}

type userRepository struct {
	db     *infra.Database
	logger *zap.Logger
}

func NewUserRepository(db *infra.Database, logger *zap.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	err := r.db.Create(&user).Error
	return errors.Wrap(err, "create user failed")
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	err := r.db.Updates(&user).Error
	return errors.Wrap(err, "update user failed")
}

func (r *userRepository) GetOneById(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, errors.Wrap(err, "get user failed")
	}

	return &user, nil
}

func (r *userRepository) GetList(ctx context.Context, req request.GetListUserRequest) ([]*model.User, *int64, error) {
	var res []*model.User
	var total int64 = 0

	query := r.db.Model(&model.User{})
	if req.Search != "" {
		query = query.Where("name like ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err := utils.QueryPagination(query, req.PageOptions, &res); err != nil {
		return nil, nil, errors.WithStack(err)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return res, &total, nil
}

func (r *userRepository) DeleteById(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return errors.Wrap(err, "Delete store failed")
	}
	return nil
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var res model.User
	err := u.db.WithContext(ctx).Where("email = ?", email).First(&res).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, err
	}
	return &res, nil
}

func (u *userRepository) GetBySocialId(ctx context.Context, socialId string) (*model.User, error) {
	var res model.User
	err := u.db.WithContext(ctx).Where("social_id = ?", socialId).First(&res).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, err
	}
	return &res, nil
}
