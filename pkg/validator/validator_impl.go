package validator

import "github.com/go-playground/validator/v10"

type TheValidatorImpl struct {
  
  Validator *validator.Validate
}

func (tvi TheValidatorImpl) Validate(i interface {}) error {

  return tvi.Validator.Struct(i);
}
