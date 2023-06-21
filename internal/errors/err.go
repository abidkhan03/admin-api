package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	err     error
	errType int
}

func New(errType int, msg string) error {
	return &Error{err: errors.New(msg), errType: errType}
}

func Newf(errType int, format string, a ...any) error {
	return &Error{err: fmt.Errorf(format, a...), errType: errType}
}

func IsNotFound(err error) bool {
	return Is(err, http.StatusNotFound)
}

func Is(err error, errType int) bool {
	gerr := &Error{}
	if errors.As(err, &gerr) {
		return gerr.errType == errType
	}
	return false
}

func Wrap(err error, errType int) error {
	return &Error{err: err, errType: errType}
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) ErrorType() int {
	return e.errType
}

const (
	InvalidArgument = iota
)

func BadRequest(msg string) error {
	return &Error{err: errors.New(msg), errType: http.StatusBadRequest}
}
