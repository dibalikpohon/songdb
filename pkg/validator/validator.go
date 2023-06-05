package validator

import "github.com/go-playground/validator/v10"

type TheValidator interface {

  Validate(i interface{}) error
}

func NewTheValidator(validator *validator.Validate) TheValidatorImpl {

  return TheValidatorImpl{ Validator: validator }
}
