package service

import (
	"context"
	"fmt"
	"log"
	"trail_backend/api/api_errors"
	"trail_backend/dto"
	dto2 "trail_backend/dto/auth"
	"trail_backend/utils/constants"

	"github.com/pkg/errors"
	"trail_backend/config"
	"trail_backend/models"

	"golang.org/x/crypto/bcrypt"
)

type (
	AuthService interface {
		Register(ctx context.Context, req dto2.RegisterRequest) (user *models.User, err error)
		Login(ctx context.Context, req dto2.LoginRequest) (res *dto2.LoginResponse, err error)
		RegisterByGoogle(ctx context.Context, req dto2.UserGoogleRequest) (user *models.User, err error)
		LoginByGoogle(ctx context.Context, req dto2.LoginByGoogleRequest) (res *dto2.LoginResponse, err error)
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

func (a *authService) Register(ctx context.Context, req dto2.RegisterRequest) (user *models.User, err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return user, err
	}

	req.Password = string(encryptedPassword)

	user, err = a.userService.Create(ctx, dto.CreateUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	return user, err
}

func (a *authService) Login(ctx context.Context, req dto2.LoginRequest) (res *dto2.LoginResponse, err error) {
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
		return nil, errors.Wrap(err, "cannot generate auth tokens")
	}

	return &dto2.LoginResponse{
		User: dto2.UserResponse{
			ID: user.ID.String(),
		},
		Token: dto2.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    a.config.Jwt.AccessTokenExpiresIn,
		},
	}, nil
}

func (a *authService) RegisterByGoogle(ctx context.Context, req dto2.UserGoogleRequest) (user *models.User, err error) {
	log.Println(fmt.Sprintf("Request info %+v", req))
	user, err = a.userService.Create(ctx, dto.CreateUserRequest{
		Email:    req.Email,
		Social:   constants.Google,
		SocialId: req.GoogleID,
	})
	return user, err
}

func (a *authService) LoginByGoogle(ctx context.Context, req dto2.LoginByGoogleRequest) (res *dto2.LoginResponse, err error) {
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
	res = &dto2.LoginResponse{
		User: dto2.UserResponse{
			ID:    user.ID.String(),
			Email: user.Email,
		},
		Token: dto2.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    a.config.Jwt.AccessTokenExpiresIn,
		},
	}

	return res, nil
}
