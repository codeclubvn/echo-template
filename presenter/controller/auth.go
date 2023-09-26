package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"trail_backend/config"
	"trail_backend/pkg/utils"
	request2 "trail_backend/presenter/request"
	"trail_backend/usecase"
)

type AuthController struct {
	BaseController
	authService usecase.AuthService
	logger      *zap.Logger
	config      *config.Config
}

func NewAuthController(authService usecase.AuthService, logger *zap.Logger, config *config.Config) *AuthController {
	controller := &AuthController{
		authService: authService,
		logger:      logger,
		config:      config,
	}
	return controller
}

// Register
//
//	@Summary		Register
//	@Description	Register
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			RegisterRequest	body		request.RegisterRequest	true	"RegisterRequest"
//	@Success		200				{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/auth/register [Post]
func (h *AuthController) Register(c echo.Context) error {
	var req request2.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	_, err := h.authService.Register(c.Request().Context(), req)

	if err != nil {
		return h.ResponseError(c, err)
	}
	return h.Response(c, http.StatusOK, "success", nil)
}

// Login
//
//	@Summary		Login
//	@Description	Login
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			LoginRequest	body		request.LoginRequest	true	"LoginRequest"
//	@Success		200				{object}	entity.LoginResponse	"success"
//	@Router			/v1/api/auth/login [Post]
func (h *AuthController) Login(c echo.Context) error {
	var req request2.LoginRequest

	if err := c.Bind(&req); err != nil {
		return h.ResponseValidationError(c, err)
	}

	res, err := h.authService.Login(c.Request().Context(), req)

	if err != nil {
		return h.ResponseError(c, err)

	}
	return h.Response(c, http.StatusOK, "success", res)
}

// GoogleLogin
//
//	@Summary		GoogleLogin
//	@Description	GoogleLogin
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		302	{object}	string
//	@Router			/v1/api/auth/google/login [Post]
func (h *AuthController) GoogleLogin(c echo.Context) error {
	authConfig := h.getGoogleOAuthConfig()
	url := authConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusFound, url)
}

func (h *AuthController) getGoogleOAuthConfig() oauth2.Config {
	return oauth2.Config{
		RedirectURL:  h.config.GoogleOAuth.RedirectURL, // Replace with your callback URL
		ClientID:     h.config.GoogleOAuth.ClientID,
		ClientSecret: h.config.GoogleOAuth.ClientSecret,
		Scopes:       h.config.GoogleOAuth.Scopes,
		Endpoint:     google.Endpoint,
	}
}

// GoogleCallback
//
//	@Summary		GoogleCallback
//	@Description	GoogleCallback
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entity.SimpleResponse	"success"
//	@Router			/v1/api/auth/call-back [Post]
func (h *AuthController) GoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	authConfig := h.getGoogleOAuthConfig()
	token, err := authConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Cannot register with google: %+v", err))
		return h.ResponseError(c, err)
	}

	client := authConfig.Client(c.Request().Context(), token)

	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token" + token.AccessToken)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Cannot register with google: %+v", err))
		return h.Response(c, http.StatusInternalServerError, "Cannot login by google", nil)
	}

	defer userInfo.Body.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(userInfo.Body)
	if err := decoder.Decode(&data); err != nil {
		// Handle JSON decoding error
		h.logger.Error(fmt.Sprintf("Cannot register with google: %+v", err))
		return h.Response(c, http.StatusInternalServerError, "Cannot login by google", nil)
	}
	//fmt.Println(data)
	var response request2.UserGoogleRequest
	if err := mapstructure.Decode(data, &response); err != nil {
		// Handle JSON unmarshaling error
		h.logger.Error(fmt.Sprintf("Cannot unmarshal JSON response: %+v", err))
		return h.Response(c, http.StatusInternalServerError, "Cannot login by google", nil)
	}

	var req request2.UserGoogleRequest

	err = utils.Copy(&response, &req)
	if err != nil {
		h.logger.Error("Cannot register with google")
		return h.Response(c, http.StatusInternalServerError, "Cannot login by google", nil)
	}
	_, err = h.authService.RegisterByGoogle(c.Request().Context(), req)
	if err != nil {
		res, err := h.authService.LoginByGoogle(c.Request().Context(), request2.LoginByGoogleRequest{
			Email:    req.Email,
			GoogleId: req.GoogleID,
		})
		if err != nil {
			return h.Response(c, http.StatusInternalServerError, "Cannot login by google account", nil)
		}

		return h.Response(c, http.StatusOK, "success", res)
	}
	return h.Response(c, http.StatusOK, "success", nil)
}
