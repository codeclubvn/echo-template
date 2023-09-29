package middlewares

import (
	"echo_template/domain/entity"
	"echo_template/pkg/api_errors"
	"echo_template/pkg/constants"
	"echo_template/presenter/request"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func (e *Middleware) Auth(authorization bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")

			if auth == "" {
				return c.JSON(http.StatusUnauthorized, entity.ResponseError{
					Message: "Unauthorized",
					Code:    api_errors.ErrUnauthorizedAccess,
				})

			}
			jwtTokenTmp := strings.Split(auth, " ")
			if len(jwtTokenTmp) < constants.NumberOfJWTTokenArray {
				mas := api_errors.MapErrorCodeMessage[api_errors.ErrTokenMissing]
				return c.JSON(mas.Status, entity.ResponseError{
					Message: mas.Message,
					Code:    api_errors.ErrTokenMissing,
				})
			}
			jwtToken := jwtTokenTmp[1]

			if jwtToken == "" {
				mas := api_errors.MapErrorCodeMessage[api_errors.ErrTokenMissing]
				return c.JSON(mas.Status, entity.ResponseError{
					Message: mas.Message,
					Code:    api_errors.ErrTokenMissing,
				})
			}

			claims, err := parseToken(jwtToken, e.config.Jwt.Secret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, entity.ResponseError{
					Message: err.Error(),
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

func parseToken(jwtToken string, secret string) (*request.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &request.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if (err.(*jwt.ValidationError)).Errors == jwt.ValidationErrorExpired {
			return nil, errors.New(api_errors.ErrTokenExpired)
		}
		return nil, errors.Wrap(err, "cannot parse token")
	}

	if claims, OK := token.Claims.(*request.JwtClaims); OK && token.Valid {
		return claims, nil
	}

	return nil, errors.New(api_errors.ErrTokenInvalid)
}
