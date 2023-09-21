package service

import (
	"context"
	"trail_backend/dto"
	"trail_backend/infrastructure"
	"trail_backend/models"
	"trail_backend/repository"
	"trail_backend/utils"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type UserService interface {
	Create(ctx context.Context, req dto.CreateUserRequest) (*models.User, error)
	UpdateUser(ctx context.Context, userID string, req dto.UpdateUserRequest) (*models.User, error)
	ListUser(ctx context.Context, req dto.ListUserRequest) ([]*models.User, *int64, error)
	DeleteUser(ctx context.Context, userID string) error
}

type usersService struct {
	userRepo repository.UserRepository
	db       *infrastructure.Database
	logger   *zap.Logger
}

func NewUserService(userRepo repository.UserRepository, db *infrastructure.Database, logger *zap.Logger) UserService {
	return &usersService{
		userRepo: userRepo,
		db:       db,
		logger:   logger,
	}
}

func (p *usersService) Create(ctx context.Context, req dto.CreateUserRequest) (*models.User, error) {
	user := &models.User{}
	if err := utils.Copy(user, req); err != nil {
		return nil, err
	}

	user, err := p.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *usersService) UpdateUser(ctx context.Context, userID string, req dto.UpdateUserRequest) (*models.User, error) {
	u := &models.User{
		Name: req.Name,
	}

	u.ID = uuid.FromStringOrNil(userID)
	user, err := p.userRepo.Update(nil, ctx, u)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *usersService) ListUser(ctx context.Context, req dto.ListUserRequest) ([]*models.User, *int64, error) {
	userID := utils.GetUserStringIDFromContext(ctx)
	users, total, err := p.userRepo.List(ctx, req.Search, req.PageOptions, userID)
	if err != nil {
		return nil, nil, err
	}

	return users, total, nil
}

func (p *usersService) DeleteUser(ctx context.Context, userID string) error {
	return p.userRepo.DeleteByID(nil, ctx, userID)
}
