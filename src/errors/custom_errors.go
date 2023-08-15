package custom_errors

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal server error")
	// Define more error types as needed
)
