package usecase

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
	"trial_backend/config"
	"trial_backend/pkg/api_errors"
	"trial_backend/pkg/constants"
	"trial_backend/presenter/request"
)

type JwtService interface {
	GenerateToken(userID string, tokenType constants.TokenType, expiresIn int64) (string, error)
	ValidateToken(token string, tokenType constants.TokenType) (*string, error)
	GenerateAuthTokens(userID string) (string, string, error)
}

type jwtService struct {
	config *config.Config
	logger *zap.Logger
}

func NewJwtService(config *config.Config, logger *zap.Logger) JwtService {
	return &jwtService{
		config: config,
		logger: logger,
	}
}

func (j *jwtService) GenerateToken(userID string, tokenType constants.TokenType, expiresIn int64) (string, error) {
	claims := request.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "trial_be",
		},
		TokenType: string(tokenType),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.config.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (j *jwtService) GenerateAuthTokens(userID string) (string, string, error) {

	j.logger.Debug("Generating request tokens", zap.Any("ExpiresIn", j.config.Jwt.AccessTokenExpiresIn))

	accessToken, err := j.GenerateToken(userID, constants.AccessToken, j.config.Jwt.AccessTokenExpiresIn)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := j.GenerateToken(userID, constants.RefreshToken, j.config.Jwt.RefreshTokenExpiresIn)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (j *jwtService) ValidateToken(token string, tokenType constants.TokenType) (*string, error) {
	claims := jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New(api_errors.ErrTokenExpired)
	}

	if claims.Issuer != "trial" {
		return nil, errors.New(api_errors.ErrTokenInvalid)
	}

	if claims.Subject == "" {
		return nil, errors.New(api_errors.ErrTokenInvalid)
	}

	return &claims.Subject, nil
}
