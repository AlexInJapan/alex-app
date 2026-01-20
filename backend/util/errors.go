package util

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidInput    = errors.New("invalid Input")
	ErrNotFound        = errors.New("not found")
	ErrInternal        = errors.New("internal server error")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrBadRequest      = errors.New("bad request")
	ErrConflict        = errors.New("conflict")
	ErrTooManyRequests = errors.New("too many requests")
	ErrNotImplemented  = errors.New("not implemented")
)

// AppError represents application-specific errors

type AppError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
	Internal   error  `json:"-"`
	StatusCode int    `json:"-"`
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Internal)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Internal
}

func NewUnauthorizedError(resource string) *AppError {
	return &AppError{
		Code:       40401,
		Message:    fmt.Sprintf("%s not found", resource),
		StatusCode: http.StatusUnauthorized,
	}
}

func NewForbiddenError() *AppError {
	return &AppError{
		Code:       40301,
		Message:    "Forbidden",
		StatusCode: http.StatusForbidden,
	}
}

func NewDatabaseError(err error) *AppError {
	return &AppError{
		Code:       50001,
		Message:    "Database operation failed",
		Internal:   err,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewInternalServerError(err error) *AppError {
	return &AppError{
		Code:       50002,
		Message:    "Internal server error",
		Internal:   err,
		StatusCode: http.StatusInternalServerError,
	}
}
