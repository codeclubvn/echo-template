package lib

import (
	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{validator: validator.New()}
}

func (cv *RequestValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
