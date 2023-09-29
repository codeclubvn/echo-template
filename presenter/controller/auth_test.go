package controller

import (
	"echo_template/pkg/lib"
	"echo_template/presenter/request"
	"echo_template/usecase/mocks"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthController_Login(t *testing.T) {
	t.Helper()
	cases := []struct {
		name    string
		req     request.LoginRequest
		want    interface{}
		wantErr int
	}{
		{
			name: "happy flow",
			req: request.LoginRequest{
				Email:    "hieuhoccode@gmail.com",
				Password: "hieuhoccode",
			},
			wantErr: 200,
		},
		{
			name: "email invalid",
			req: request.LoginRequest{
				Email:    "hieuhoccodil.com",
				Password: "hieuhoccode",
			},
			wantErr: 422,
		},
		{
			name: "email is missing",
			req: request.LoginRequest{
				Email:    "",
				Password: "hieuhoccode",
			},
			wantErr: 422,
		},
		{
			name: "password is blank space",
			req: request.LoginRequest{
				Email:    "hieuhoccode@gmail.com",
				Password: "",
			},
			wantErr: 422,
		},
	}

	// Setup
	e := echo.New()
	e.Validator = lib.NewRequestValidator()

	mockUCase := new(mocks.AuthService)
	h := &AuthController{
		authService: mockUCase,
	}

	// Mocking
	mockUCase.On("Login", echo.POST, mock.Anything).Return(nil, nil)

	// Gọi hàm Login của tầng handler
	for _, item := range cases {
		b, err := json.Marshal(item.req)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/v1/api/auth/login", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h.Login(c)) {
			assert.Equal(t, item.wantErr, rec.Code)
		}
	}

	// Kiểm tra rằng authService.Login đã được gọi một lần
	mockUCase.AssertCalled(t, "Login", mock.Anything, mock.Anything)
}

func TestAuthController_Register(t *testing.T) {
	cases := []struct {
		name    string
		req     request.RegisterRequest
		want    interface{}
		wantErr int
	}{
		{
			name: "happy flow",
			req: request.RegisterRequest{
				Email:    "hieuhoccode@gmail.com",
				Password: "hieuhoccode",
			},
			wantErr: 200,
		},
		{
			name: "email invalid",
			req: request.RegisterRequest{
				Email:    "hieuhoccodil.com",
				Password: "hieuhoccode",
			},
			wantErr: 422,
		},
		{
			name: "email is missing",
			req: request.RegisterRequest{
				Email:    "",
				Password: "hieuhoccode",
			},
			wantErr: 422,
		},
		{
			name: "password is blank space",
			req: request.RegisterRequest{
				Email:    "hieuhoccode@gmail.com",
				Password: "",
			},
			wantErr: 422,
		},
	}

	// Setup
	e := echo.New()
	e.Validator = lib.NewRequestValidator()

	mockUCase := new(mocks.AuthService)
	h := &AuthController{
		authService: mockUCase,
	}

	// Mocking
	mockUCase.On("Register", mock.Anything, mock.Anything).Return(nil, nil)

	// Gọi hàm Register của tầng handler
	for _, item := range cases {
		b, err := json.Marshal(item.req)
		if err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/v1/api/auth/register", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h.Register(c)) {
			assert.Equal(t, item.wantErr, rec.Code)
		}
	}

	// Kiểm tra rằng authService.Login đã được gọi một lần
	mockUCase.AssertCalled(t, "Register", mock.Anything, mock.Anything)
}
