package utils

import (
	"errors"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"trial_backend/pkg/api_errors"
)

func GetUserStringIDFromContext(ctx echo.Context) string {
	return ctx.Request().Header.Get("x-user-id")
}

func GetUserUUIDFromContext(ctx echo.Context) (uuid.UUID, error) {
	sid := ctx.Request().Header.Get("x-user-id")

	u, err := uuid.FromString(sid)
	if err != nil {
		return uuid.Nil, errors.New(api_errors.ErrInvalidUserID)
	}

	return u, nil
}

func GetPageCount(total int64, limit int64) int64 {
	if total == 0 {
		return 0
	}

	if total%limit != 0 {
		return total/limit + 1
	}

	return total / limit
}

func ParseStringIDFromUri(c echo.Context) string {
	id := c.Param("id")
	if len(id) == 0 {
		return ""
	}
	return id
}
