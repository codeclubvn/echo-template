package repository

import (
	"context"
	"time"
	"trail_backend/api/api_errors"
	"trail_backend/api/dto"
	"trail_backend/infrastructure"
	"trail_backend/models"
	"trail_backend/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (err error)
	Update(ctx context.Context, user *models.User) (err error)
	GetList(ctx context.Context, req dto.GetListUserRequest) (res *dto.ListUserResponse, total *int64, err error)
	GetOneById(ctx context.Context, req dto.GetOneUserRequest) (res *models.User, err error)
	DeleteById(ctx context.Context, req dto.DeleteUserRequest) (err error)
	GetByEmail(ctx context.Context, email string) (res *models.User, err error)
	GetBySocialId(ctx context.Context, socialId string) (res *models.User, err error)
}

type userRepository struct {
	db     *infrastructure.Database
	logger *zap.Logger
}

func NewUserRepository(db *infrastructure.Database, logger *zap.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) (err error) {
	err = r.db.Create(&user).Error
	return errors.Wrap(err, "create user failed")
}

func (r *userRepository) Update(ctx context.Context, user *models.User) (err error) {
	err = r.db.Updates(&user).Error
	return errors.Wrap(err, "update user failed")
}

func (r *userRepository) GetOneById(ctx context.Context, req dto.GetOneUserRequest) (res *models.User, err error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("id = ?", req.Id).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "10006")
	}

	return &user, nil
}

func (r *userRepository) GetList(ctx context.Context, req dto.GetListUserRequest) (res *dto.ListUserResponse, total *int64, err error) {
	query := r.db.Model(&models.User{})
	if req.Search != "" {
		query = query.Where("name like ?", "%"+req.Search+"%")
	}

	switch req.Sort {
	default:
		query = query.Order(req.Sort)
	}

	if err = utils.QueryPagination(query, req.PageOptions, &res.Data); err != nil {
		return nil, nil, errors.WithStack(err)
	}

	if err = query.Count(total).Error; err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return res, total, err
}

func (r *userRepository) DeleteById(ctx context.Context, req dto.DeleteUserRequest) (err error) {
	err = r.db.Where("id = ?", req.Id).Updates(map[string]interface{}{"deleted_at": time.Time{}, "updater_id": req.UserId}).Error
	return errors.Wrap(err, "delete user failed")
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (res *models.User, err error) {
	err = u.db.WithContext(ctx).Where("email = ?", email).First(&res).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, err
	}
	return
}

func (u *userRepository) GetBySocialId(ctx context.Context, socialId string) (res *models.User, err error) {
	err = u.db.WithContext(ctx).Where("social_id = ?", socialId).First(&res).Error
	if err != nil {
		if utils.ErrNoRows(err) {
			return nil, errors.New(api_errors.ErrUserNotFound)
		}
		return nil, err
	}
	return
}
