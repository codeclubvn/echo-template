package usecase

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"trail_backend/domain/repo"
	"trail_backend/domain/repo/model"
	"trail_backend/infra"
	"trail_backend/pkg/utils"
	"trail_backend/presenter/request"
)

type UserService interface {
	Create(ctx context.Context, req request.CreateUserRequest) (*model.User, error)
	Update(ctx context.Context, req request.UpdateUserRequest) (*model.User, error)
	GetList(ctx context.Context, req request.GetListUserRequest) ([]*model.User, *int64, error)
	Delete(ctx context.Context, id string) error
	GetOne(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetBySocialId(ctx context.Context, socialId string) (*model.User, error)
}

type userService struct {
	userRepository       repo.UserRepository
	cloudinaryRepository repo.CloudinaryRepository
	db                   *infra.Database
	logger               *zap.Logger
}

func NewUserService(userRepo repo.UserRepository, cloudinaryRepository repo.CloudinaryRepository, db *infra.Database, logger *zap.Logger) UserService {
	return &userService{
		userRepository:       userRepo,
		cloudinaryRepository: cloudinaryRepository,
		db:                   db,
		logger:               logger,
	}
}

func (s *userService) Create(ctx context.Context, req request.CreateUserRequest) (*model.User, error) {
	user := &model.User{}
	if err := utils.Copy(user, req); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Update(ctx context.Context, req request.UpdateUserRequest) (*model.User, error) {
	user, err := s.userRepository.GetOneById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if err = utils.Copy(user, req); err != nil {
		return nil, err
	}
	if err = s.userRepository.Update(ctx, user); err != nil {
		return nil, err
	}
	return user, err
}

func (s *userService) GetList(ctx context.Context, req request.GetListUserRequest) ([]*model.User, *int64, error) {
	return s.userRepository.GetList(ctx, req)
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return s.userRepository.DeleteById(ctx, id)
}

func (s *userService) GetOne(ctx context.Context, id string) (*model.User, error) {
	return s.userRepository.GetOneById(ctx, id)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.userRepository.GetByEmail(ctx, email)
}

func (s *userService) GetBySocialId(ctx context.Context, socialId string) (*model.User, error) {
	return s.userRepository.GetBySocialId(ctx, socialId)
}
