package errors

import (
	"errors"
	"fmt"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExist = errors.New("user already exists")
	ErrInvalidUserID    = errors.New("invalid user id")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrDatabaseError    = errors.New("database error")
	ErrCacheError       = errors.New("cache error")
	ErrSearchError      = errors.New("search error")
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}
