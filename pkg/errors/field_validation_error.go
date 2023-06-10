package errors

import "fmt"

type fieldError struct {
  FieldName string `json:"fieldName"`
  Why string `json:"why"`
}

type FieldValidationError struct {
  Message string `json:"message"`
  Errors []fieldError `json:"errors"`
}

func NewFieldValidationError() FieldValidationError {
  
  return FieldValidationError{Message: "Field validation error"}
}

func (fve *FieldValidationError) AppendError(fieldName string, why string) {
  
  fve.Errors = append(fve.Errors, fieldError{fieldName, why})
}

func (fe *fieldError) Error() string {

  return fmt.Sprintf("('%s': %s)", fe.FieldName, fe.Why)
}

func (fve *FieldValidationError) Error() string {

  return fmt.Sprintf("Field validation error. %s", fmt.Sprintf("%+q", fve.Errors))
}
