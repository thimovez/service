package validator

import (
	"github.com/go-playground/validator/v10"
)

type IValidator interface {
	ValidateStruct(s interface{}) error
}

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	validate := validator.New()

	v := &Validator{validate: validate}

	return v
}

func (v *Validator) ValidateStruct(s interface{}) error {
	err := v.validate.Struct(s)
	if err != nil {
		err = err.(validator.ValidationErrors)
		//fmt.Sprintf("Validation error: %s", errors)
		return err
	}

	return nil
}
