package cerror

import (
	"fmt"

	"github.com/pkg/errors"
)

// errorCode представляет собой код ошибки JSONRPC
type errorCode int

const (
	Unknown errorCode = 100 + iota
	NotFound
	BadRequest
)

// CustomError выдает код и сообщение об ошибке
type CustomError interface {
	Code() int
	Error() string
}

// customError имплементирует CustomError
type customError struct {
	code errorCode
	err  error
}

func (e customError) Code() int {
	return int(e.code)
}

func (e customError) Error() string {
	return e.err.Error()
}

func (e customError) Cause() error {
	return e.err
}

func (code errorCode) New(msg string) error {
	return customError{code: code, err: errors.New(msg)}
}

func (code errorCode) Newf(msg string, args ...interface{}) error {
	return customError{code: code, err: fmt.Errorf(msg, args...)}
}

func (code errorCode) Wrap(err error, msg string) error {
	return code.Wrapf(err, msg)
}

func (code errorCode) Wrapf(err error, msg string, args ...interface{}) error {
	return customError{code: code, err: errors.Wrapf(err, msg, args...)}
}

// func WithMessage(err error, msg string) error {
// 	we := errors.WithMessage(err, msg)

// 	if ce, ok := err.(customError); ok {
// 		return customError{code: ce.code, err: we}
// 	}

// 	return customError{code: Unknown, err: we}
// }
