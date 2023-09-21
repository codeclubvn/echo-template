package utils

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"trail_backend/api/api_errors"

	uuid "github.com/satori/go.uuid"
)

func GetUserStringIDFromContext(ctx context.Context) string {
	return ctx.Value("x-user-id").(string)
}

func GetUserUUIDFromContext(ctx context.Context) (uuid.UUID, error) {
	sid := ctx.Value("x-user-id").(string)

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
