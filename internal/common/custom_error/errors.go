package custom_error

import "errors"

type ErrorCode int

const (
	ErrNotFound ErrorCode = iota
	ErrInternal
	ErrInvalidInput
	ErrConflict
)

type Error struct {
	message string
	err     error
	code    ErrorCode
}

func (e *Error) Error() string { return e.err.Error() }

func (e *Error) Unwrap() error { return e.err }

func (e *Error) Code() ErrorCode { return e.code }

func (e *Error) Message() string { return e.message }

func NewError(err error, code ErrorCode, message string) *Error {
	return &Error{
		message: message,
		err:     err,
		code:    code,
	}
}

func NewBadInputError(err error, m ...string) *Error {
	message := "invalid input parameters"
	if len(m) > 0 {
		message = m[0]
	}
	return NewError(err, ErrInvalidInput, message)
}

func NewInternalError(err error, m ...string) *Error {
	message := "internal server error"
	if len(m) > 0 {
		message = m[0]
	}
	return NewError(err, ErrInternal, message)
}

func NewNotFoundError(err error, m ...string) *Error {
	message := "resource not found"
	if len(m) > 0 {
		message = m[0]
	}
	return NewError(err, ErrNotFound, message)
}

func NewConflictError(err error, m ...string) *Error {
	message := "conflict occured"
	if len(m) > 0 {
		message = m[0]
	}
	return NewError(err, ErrConflict, message)
}

func IsCustom(err error) (*Error, bool) {
	var c *Error
	return c, errors.As(err, &c)
}
