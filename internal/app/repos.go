package app

import (
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app/config"
)

type Repos struct {
	TeacherRepo internal.TeacherRepo
	StudentRepo internal.StudentRepo
	WebinarRepo internal.WebinarRepo
}

func newRepos(cfg *config.Config) (*Repos, error) {
	return &Repos{
	}, nil
}

func (*Repos) Close() []error {
	var errs []error
	return errs
}
