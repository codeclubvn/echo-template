package usecase

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"trial_backend/config"
	"trial_backend/domain/entity"
	"trial_backend/domain/repo/model"
	"trial_backend/pkg/api_errors"
	"trial_backend/pkg/constants"
	"trial_backend/presenter/request"
)

type (
	AuthService interface {
		Register(ctx context.Context, req request.RegisterRequest) (user *model.User, err error)
		Login(ctx context.Context, req request.LoginRequest) (res *entity.LoginResponse, err error)
		RegisterByGoogle(ctx context.Context, req request.UserGoogleRequest) (user *model.User, err error)
		LoginByGoogle(ctx context.Context, req request.LoginByGoogleRequest) (res *entity.LoginResponse, err error)
	}
	authService struct {
		userService UserService
		jwtService  JwtService
		config      *config.Config
	}
)

func NewAuthService(userService UserService, config *config.Config, jwtService JwtService) AuthService {
	return &authService{
		userService: userService,
		jwtService:  jwtService,
		config:      config,
	}
}

func (a *authService) Register(ctx context.Context, req request.RegisterRequest) (user *model.User, err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return user, err
	}

	req.Password = string(encryptedPassword)

	user, err = a.userService.Create(ctx, request.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	return user, err
}

func (a *authService) Login(ctx context.Context, req request.LoginRequest) (res *entity.LoginResponse, err error) {
	user, err := a.userService.GetByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, errors.New(api_errors.ErrInvalidPassword)
	}

	accessToken, refreshToken, err := a.jwtService.GenerateAuthTokens(user.ID.String())
	if err != nil {
		return nil, errors.Wrap(err, "cannot generate request tokens")
	}

	return &entity.LoginResponse{
		User: entity.UserResponse{
			ID: user.ID.String(),
		},
		Token: entity.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (a *authService) RegisterByGoogle(ctx context.Context, req request.UserGoogleRequest) (user *model.User, err error) {
	log.Println(fmt.Sprintf("Request info %+v", req))
	user, err = a.userService.Create(ctx, request.CreateUserRequest{
		Email:    req.Email,
		Social:   constants.Google,
		SocialId: req.GoogleID,
	})
	return user, err
}

func (a *authService) LoginByGoogle(ctx context.Context, req request.LoginByGoogleRequest) (res *entity.LoginResponse, err error) {
	user, err := a.userService.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	_, err = a.userService.GetBySocialId(ctx, req.GoogleId)
	if err != nil {
		return nil, err
	}
	accessToken, refreshToken, err := a.jwtService.GenerateAuthTokens(user.ID.String())
	if err != nil {
		return nil, err
	}
	res = &entity.LoginResponse{
		User: entity.UserResponse{
			ID:    user.ID.String(),
			Email: user.Email,
		},
		Token: entity.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return res, nil
}
