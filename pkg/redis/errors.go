package redis

import "errors"

var (
	ErrHasOpenSession      = errors.New("user has an open session")
	ErrSessionNotExists    = errors.New("session does not exists")
	ErrInvalidSessionToken = errors.New("invalid session token")
	ErrUserNotExists       = errors.New("user does not exists")
	ErrNotAuthorized       = errors.New("not authorized")
)
