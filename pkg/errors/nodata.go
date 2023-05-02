package errors

type NoData struct {
  Message string    `json:"message"`   // Error message
  What interface{}  `json:"what"`     // What the caller is trying to look at
}

func (e *NoData) Error() string {
  return e.Message
}
