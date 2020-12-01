package entity

import "github.com/pkg/errors"

var (
	ErrUserNotFound        = errors.New("user is not found")
	ErrIDNotFoundInSession = errors.New("id is not found in session")
	//ErrFailedToCreateWebinar = errors.New("failed creating new webinar")
)
