package errors

import "errors"

// General Errors
var (
	ErrInvalidIpFormat = errors.New("invalid ip format")
	ErrBannedUser      = errors.New("user banned")
	ErrUserDBNotFound  = errors.New("user not found in DB")
	ErrUserNotAdmin    = errors.New("user is not an admin and cannot ban others")
)
