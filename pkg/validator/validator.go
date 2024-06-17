package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type IValidator interface {
	ValidateStruct(s interface{}) (errors gin.H)
}

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	validate := validator.New()

	v := &Validator{validate: validate}

	return v
}

func (v *Validator) ValidateStruct(s interface{}) (errors gin.H) {
	err := v.validate.Struct(s)
	if err != nil {
		errors := gin.H{}
		for _, err := range err.(validator.ValidationErrors) {
			e := err.Error()
			errors[err.Field()] = e
		}
		return errors
	}

	return nil
}
