package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"trial_backend/pkg/lib"
	"trial_backend/presenter/request"
)

func TestAuthController_Login(t *testing.T) {

	cases := request.LoginRequest{
		Email:    "hieuhoccode@gmail.com",
		Password: "hieuhoccode",
	}

	b, err := json.Marshal(cases)
	if err != nil {
		t.Fatal(err)
	}

	// Setup
	e := echo.New()
	e.Validator = lib.NewRequestValidator()

	req := httptest.NewRequest(http.MethodPost, "/v1/api/auth/login", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &AuthController{
		BaseController: BaseController{},
	}

	// Assertions

	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

}
