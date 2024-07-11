package apperrors

import (
	"errors"
)

type RequestError struct {
	Code int   `json:"code"`
	Err  error `json:"Error"`
}

func NewAppError(code int, err error) *RequestError {
	return &RequestError{
		Code: code,
		Err:  err,
	}
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func ErrNotFound() *RequestError {
	return &RequestError{
		Code: 404,
		Err:  errors.New("404 not found"),
	}
}
func ErrInternalServer() *RequestError {
	return &RequestError{
		Code: 500,
		Err:  errors.New("500 internal server error"),
	}
}
func ErrBadRequest() *RequestError {
	return &RequestError{
		Code: 400,
		Err:  errors.New("400 bad request"),
	}
}

func ErrUnAuthorized() *RequestError {
	return &RequestError{
		Code: 401,
		Err:  errors.New("401 unauthorized"),
	}
}
func ErrForbidden() *RequestError {
	return &RequestError{
		Code: 403,
		Err:  errors.New("403 Forbidden"),
	}
}
