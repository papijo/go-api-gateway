package errorservice

import (
	"fmt"
)

// AppError represents an application-specific error with a code and message.
type AppError struct {
	Code    int
	Message string
	Details map[string]interface{}
}

// New creates a new AppError with the specified code and message.
func NewError(code int, message string, details map[string]interface{}) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// Error implements the error interface.
func (e *AppError) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

// HTTPStatus returns the HTTP status code for this error.
func (e *AppError) HTTPStatus() int {
	return e.Code
}
