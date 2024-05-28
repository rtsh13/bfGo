package errors

import (
	"time"
)

type ErrorResponse struct {
	Reason   interface{}
	DateTime time.Time
}

func (e *ErrorResponse) Error() interface{} {
	return e.Reason
}

type Error string

func (e Error) Error() string {
	return string(e)
}
