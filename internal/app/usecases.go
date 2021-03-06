package app

import (
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app/config"
	"github.com/ridwanakf/weebinar/internal/usecase"
)

type Usecases struct {
	TeacherUC internal.TeacherUC
	StudentUC internal.StudentUC
}

func newUsecases(repos *Repos, cfg config.Config) *Usecases {
	return &Usecases{
		TeacherUC: usecase.NewTeacherUsecase(repos.TeacherRepo, repos.WebinarRepo),
		StudentUC: usecase.NewStudentUsecase(repos.StudentRepo, repos.WebinarRepo),
	}
}

func (*Usecases) Close() []error {
	var errs []error
	return errs
}
