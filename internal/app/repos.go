package app

import (
	"github.com/ridwanakf/weebinar/internal/app/config"
)

type Repos struct {

}

func newRepos(cfg *config.Config) (*Repos, error) {
	return &Repos{
	}, nil
}

func (*Repos) Close() []error {
	var errs []error
	return errs
}
