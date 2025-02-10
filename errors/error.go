package errors

import "errors"

// General Errors
var (
	ErrInvalidIpFormat = errors.New("invalid ip format")
	ErrBannedUser      = errors.New("user banned")
)
