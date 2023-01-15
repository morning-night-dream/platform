package errors

import "errors"

type NotFoundError struct {
	msg string
	err error
}

func NewNotFoundError(msg string, err error) NotFoundError {
	if msg == "" && err != nil {
		msg = err.Error()
	}

	return NotFoundError{
		msg: msg,
		err: err,
	}
}

func (e NotFoundError) Error() string { return e.msg }

func (e NotFoundError) Unwrap() error { return e.err }

func AsNotFoundError(err error) bool {
	var target NotFoundError

	return errors.As(err, &target)
}

type ValidationError struct {
	msg string
	err error
}

func NewValidationError(msg string, err error) ValidationError {
	if msg == "" && err != nil {
		msg = err.Error()
	}

	return ValidationError{
		msg: msg,
		err: err,
	}
}

func (e ValidationError) Error() string { return e.msg }

func (e ValidationError) Unwrap() error { return e.err }
