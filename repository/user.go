package repository

import (
	"context"
	"trail_backend/dto"
	"trail_backend/infrastructure"
	"trail_backend/models"
	"trail_backend/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(tx *TX, ctx context.Context, user *models.User) (*models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
	List(ctx context.Context, search string, o dto.PageOptions, userID string) ([]*models.User, *int64, error)
	DeleteByID(tx *TX, ctx context.Context, id string) error
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

func (p *userRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user.UpdaterID = currentUID

	if err := p.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, errors.Wrap(err, "Create user failed")
	}

	return user, nil
}

func (p *userRepository) Update(tx *TX, ctx context.Context, user *models.User) (*models.User, error) {
	currentUID, err := utils.GetUserUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		tx = &TX{db: *p.db}
	}

	user.UpdaterID = currentUID
	if err := tx.db.WithContext(ctx).Save(user).Error; err != nil {
		return nil, errors.Wrap(err, "Update user failed")
	}

	return user, nil
}

func (p *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := p.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.Wrap(err, "Find user failed")
	}

	return &user, nil
}

func (p *userRepository) List(ctx context.Context, search string, o dto.PageOptions, userID string) ([]*models.User, *int64, error) {
	var users []*models.User
	var total int64 = 0

	q := p.db.WithContext(ctx).Model(&models.User{})

	if search != "" {
		q = q.Where("name LIKE ?", "%"+search+"%")
	}

	q.Order("created_at DESC")

	if err := utils.QueryPagination(p.db, o, &users).Count(&total).Error(); err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return users, &total, nil
}

func (p *userRepository) DeleteByID(tx *TX, ctx context.Context, id string) error {
	if tx == nil {
		tx = &TX{db: *p.db}
	}

	if err := tx.db.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return errors.Wrap(err, "Delete user failed")
	}

	return nil
}
