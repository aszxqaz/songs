package validator

import (
	"time"

	validator "github.com/go-playground/validator/v10"
)

func dmyDate(sl validator.FieldLevel) bool {
	_, err := time.Parse("02.01.2006", sl.Field().String())
	return err == nil
}
