package usecase

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"trial_backend/domain/entity"
	"trial_backend/domain/repo/mocks"
	"trial_backend/domain/repo/model"
	"trial_backend/pkg/api_errors"
	"trial_backend/presenter/request"
	mockjwt "trial_backend/usecase/mocks"
)

func Test_authService_Login(t *testing.T) {
	mockURepo := new(mocks.UserRepository)
	mockjwt := new(mockjwt.JwtService)

	cases := []struct {
		name    string
		ctx     context.Context
		req     request.LoginRequest
		want    *entity.LoginResponse
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "happy flow",
			ctx:  context.Background(),
			req: request.LoginRequest{
				Email:    "hieuhoccode@gmail.com",
				Password: "hieuhoccode",
			},
			wantErr: nil,
		},
		{
			name: "wrong password",
			ctx:  context.Background(),
			req: request.LoginRequest{
				Email:    "hieuhoccode@gmail.com",
				Password: "hieuhoccode1",
			},
			wantErr: errors.New(api_errors.ErrInvalidPassword),
		},
	}
	res := &model.User{
		BaseModel: model.BaseModel{
			ID: uuid.FromStringOrNil("e9bdd7b3-003d-46fc-bab0-5105597731ce"),
		},
		Email:    "hieuhoccode@gmail.com",
		Password: "$2a$10$/wvln13sGh82r.bGiWxdMO3JWgV8sfhoh5P1noGsi49IUSm2oSe/m",
	}
	mockURepo.On("GetByEmail", mock.Anything, "hieuhoccode@gmail.com").Return(res, nil)
	mockjwt.On("GenerateAuthTokens", mock.Anything, mock.AnythingOfType("string")).Return("token", "refresh_token", nil)

	s := &authService{
		userService: &userService{
			userRepository: mockURepo,
		},
		jwtService: mockjwt,
	}

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			_, err := s.Login(item.ctx, item.req)
			assert.Equal(t, item.wantErr, err)
		})
	}
}
