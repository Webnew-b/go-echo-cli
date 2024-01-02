package controller

import (
	"github.com/go-playground/validator/v10"
)

var valid = validator.New()

func GetValidator() *validator.Validate {
	return valid
}
