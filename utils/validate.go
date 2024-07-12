package utils

import "github.com/go-playground/validator/v10"

type BindingValidation interface {
	Validate(data interface{}) error
}
type bindingValidator struct {
	validator *validator.Validate
}

func NewValidator() *bindingValidator {
	var validate = validator.New()
	return &bindingValidator{
		validator: validate,
	}
}
func (v *bindingValidator) Validate(data interface{}) error {
	errs := v.validator.Struct(data)
	if errs != nil {
		return errs

	}

	return nil
}
