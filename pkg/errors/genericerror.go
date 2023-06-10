package errors

import "fmt"

type GenericError struct {
  Message string `json:"message"`
  AnotherMessage string `json:"anotherMessage"`
}

func (ge *GenericError) Error() string {
  return fmt.Sprintf("Generic error. Error: \"%s\". AnotherMessage: \"%s\"", ge.Message, ge.AnotherMessage);
}

func (ge *GenericError) MalformedPayload(anotherMessage string) error {
  ge.Message = "Malformed Payload"
  ge.AnotherMessage = anotherMessage

  return ge
}
