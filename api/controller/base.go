package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"trail_backend/api/api_errors"
	dto2 "trail_backend/dto"
	"trail_backend/utils"

	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

type BaseController struct {
}

func (b *BaseController) Response(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, dto2.SimpleResponse{
		Data:    data,
		Message: message,
	})
}

func (b *BaseController) ResponseList(c echo.Context, message string, total *int64, data interface{}) error {
	var o dto2.PageOptions
	if err := c.Bind(&o); err != nil {
		return b.ResponseValidationError(c, err)
	}

	if o.Limit == 0 {
		o.Limit = 10
	}

	if o.Page == 0 {
		o.Page = 1
	}

	pageCount := utils.GetPageCount(*total, o.Limit)
	return c.JSON(http.StatusOK, dto2.SimpleResponseList{
		Message: message,
		Data:    data,
		Meta: dto2.Meta{
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

	return c.JSON(status, dto2.ResponseError{
		Code:    err.Error(),
		Message: mas.Message,
	})
}

func (b *BaseController) ResponseValidationError(c echo.Context, err error) error {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		err = errors.New(utils.StructPascalToSnakeCase(ve[0].Field()) + " is " + ve[0].Tag())
	}

	return c.JSON(http.StatusUnprocessableEntity, dto2.ResponseError{
		Code:    api_errors.ErrValidation,
		Message: err.Error(),
	})
}
