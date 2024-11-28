package validator

import validator "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
	Validate.RegisterValidation("dmyDate", dmyDate)
}
