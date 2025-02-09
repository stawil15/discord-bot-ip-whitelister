package errors

import "errors"

// General Errors
var (
	ErrInvalidIpFormat = errors.New("invalid ip format")
)

// Database Errors
var (
	ErrDBConnectionFailed = errors.New("database connection failed")
	ErrDBTimeout          = errors.New("database connection timeout")
)
