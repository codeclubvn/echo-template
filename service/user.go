package service

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"trail_backend/api/dto"
	"trail_backend/infrastructure"
	"trail_backend/models"
	"trail_backend/repository"
	"trail_backend/utils"

	uuid "github.com/satori/go.uuid"
)

type UserService interface {
	Create(ctx context.Context, req dto.CreateUserRequest) (*models.User, error)
	Update(ctx context.Context, req dto.UpdateUserRequest) (*models.User, error)
	GetList(ctx context.Context, req dto.GetListUserRequest) (*dto.ListUserResponse, error)
	Delete(ctx context.Context, req dto.DeleteUserRequest) error
	GetOne(ctx context.Context, req dto.GetOneUserRequest) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetBySocialId(ctx context.Context, socialId string) (*models.User, error)
}

type userService struct {
	userRepository       repository.UserRepository
	cloudinaryRepository repository.CloudinaryRepository
	db                   *infrastructure.Database
	logger               *zap.Logger
}

func NewUserService(userRepo repository.UserRepository, cloudinaryRepository repository.CloudinaryRepository, db *infrastructure.Database, logger *zap.Logger) UserService {
	return &userService{
		userRepository:       userRepo,
		cloudinaryRepository: cloudinaryRepository,
		db:                   db,
		logger:               logger,
	}
}

func (s *userService) Create(ctx context.Context, req dto.CreateUserRequest) (*models.User, error) {
	user := &models.User{}
	var err error

	if err = utils.Copy(user, req); err != nil {
		return nil, errors.WithStack(err)
	}

	if err = s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, err
}

func (s *userService) Update(ctx context.Context, req dto.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepository.GetOneById(ctx, dto.GetOneUserRequest{Id: req.UserId})
	if err != nil {
		return nil, err
	}

	if err = utils.Copy(user, req); err != nil {
		return nil, err
	}
	user.UpdaterID = uuid.FromStringOrNil(req.UserId)

	if err = s.userRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) GetList(ctx context.Context, req dto.GetListUserRequest) (*dto.ListUserResponse, error) {
	res := &dto.ListUserResponse{}
	var err error

	res, err = s.userRepository.GetList(ctx, req)
	return res, err
}

func (s *userService) Delete(ctx context.Context, req dto.DeleteUserRequest) error {
	err := s.userRepository.DeleteById(ctx, req)
	return err
}

func (s *userService) GetOne(ctx context.Context, req dto.GetOneUserRequest) (*models.User, error) {
	user := &models.User{}
	var err error

	user, err = s.userRepository.GetOneById(ctx, req)
	return user, err
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, err
		}
		return nil, err
	}
	return user, err
}

func (s *userService) GetBySocialId(ctx context.Context, socialId string) (*models.User, error) {
	user, err := s.userRepository.GetBySocialId(ctx, socialId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
