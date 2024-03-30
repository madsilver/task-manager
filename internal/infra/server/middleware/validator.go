package middleware

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func ConfigValidator() *CustomValidator {
	return &CustomValidator{
		Validator: validator.New(),
	}
}

func (c *CustomValidator) Validate(i any) error {
	return c.Validator.Struct(i)
}
