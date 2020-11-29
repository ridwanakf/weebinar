package app

import (
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app/config"
)

type Usecases struct {
	TeacherUC internal.TeacherUC
	StudentUC internal.StudentUC
}

func newUsecases(repos *Repos, gateway *Gateways, cfg *config.Config) *Usecases {
	return &Usecases{
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
