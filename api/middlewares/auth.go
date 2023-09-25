package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"trail_backend/api/api_errors"
	"trail_backend/dto"
	dtoAuth "trail_backend/dto/auth"

	"github.com/pkg/errors"

	"github.com/golang-jwt/jwt"
)

func (e *Middleware) Auth(authorization bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")

			if auth == "" {
				return c.JSON(http.StatusUnauthorized, dto.ResponseError{
					Message: "Unauthorized",
					Code:    api_errors.ErrUnauthorizedAccess,
				})

			}
			jwtToken := strings.Split(auth, " ")[1]

			if jwtToken == "" {
				mas := api_errors.MapErrorCodeMessage[api_errors.ErrTokenMissing]
				return c.JSON(mas.Status, dto.ResponseError{
					Message: mas.Message,
					Code:    api_errors.ErrTokenMissing,
				})
			}

			claims, err := parseToken(jwtToken, e.config.Jwt.Secret)
			if err != nil {
				mas := api_errors.MapErrorCodeMessage[err.Error()]
				return c.JSON(mas.Status, dto.ResponseError{
					Message: mas.Message,
					Code:    api_errors.ErrTokenInvalid,
				})
			}

			c.Request().Header.Add("x-user-id", claims.Subject)
			if !authorization {
				return next(c)
			}

			return next(c)
		}
	}
}

func parseToken(jwtToken string, secret string) (*dtoAuth.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &dtoAuth.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if (err.(*jwt.ValidationError)).Errors == jwt.ValidationErrorExpired {
			return nil, errors.New(api_errors.ErrTokenExpired)
		}
		return nil, errors.Wrap(err, "cannot parse token")
	}

	if claims, OK := token.Claims.(*dtoAuth.JwtClaims); OK && token.Valid {
		return claims, nil
	}

	return nil, errors.New(api_errors.ErrTokenInvalid)
}
