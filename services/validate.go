package services

import (
	"github.com/go-playground/validator"
)

func ValidateForm(i interface{}) (bool, error) {
	v := validator.New()
	_ = v.RegisterValidation("customoneof", func(fl validator.FieldLevel) bool {
		return (fl.Field().String()) == "Let´s Work On This" || (fl.Field().String()) == "Reach The Goal" || (fl.Field().String()) == "Relevant Performance" || (fl.Field().String()) == "Master"
	})
	err := v.Struct(i)
	if err != nil {
		return false, err
	}
	return true, nil
}
