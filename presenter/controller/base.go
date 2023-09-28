package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"trial_backend/domain/entity"
	"trial_backend/pkg/api_errors"
	"trial_backend/pkg/utils"
	"trial_backend/presenter/request"
)

type BaseController struct {
}

func (b *BaseController) Response(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, entity.SimpleResponse{
		Data:    data,
		Message: message,
	})
}

func (b *BaseController) ResponseList(c echo.Context, message string, total *int64, data interface{}) error {
	var o request.PageOptions
	if err := c.Bind(&o); err != nil {
		return b.ResponseValidatorError(c, err)
	}

	if o.Limit == 0 {
		o.Limit = 10
	}

	if o.Page == 0 {
		o.Page = 1
	}

	pageCount := utils.GetPageCount(*total, o.Limit)
	return c.JSON(http.StatusOK, entity.SimpleResponseList{
		Message: message,
		Data:    data,
		Meta: entity.Meta{
			Total:       total,
			Page:        o.Page,
			Limit:       o.Limit,
			Sort:        o.Sort,
			PageCount:   pageCount,
			HasPrevPage: o.Page > 1,
			HasNextPage: o.Page < pageCount,
		},
	})
}

func (b *BaseController) ResponseError(c echo.Context, err error) error {
	fmt.Println(err.Error())
	mas, ok := api_errors.MapErrorCodeMessage[err.Error()]
	status := mas.Status
	if !ok {
		status = http.StatusInternalServerError
		mas = api_errors.MapErrorCodeMessage[api_errors.ErrInternalServerError]
		err = errors.New(api_errors.ErrInternalServerError)
	}

	return c.JSON(status, entity.ResponseError{
		Code:    err.Error(),
		Message: mas.Message,
	})
}

func (b *BaseController) ResponseValidatorError(c echo.Context, err error) error {
	message := strings.Split(err.Error(), "message=")[0]

	return c.JSON(http.StatusUnprocessableEntity, entity.ResponseError{
		Code:    api_errors.ErrValidation,
		Message: message,
	})
}
