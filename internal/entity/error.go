package entity

import "github.com/pkg/errors"

var (
	ErrUserNotFound = errors.New("user is not found")
	//ErrFailedToCreateWebinar = errors.New("failed creating new webinar")
)

