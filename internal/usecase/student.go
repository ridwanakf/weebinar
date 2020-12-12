package usecase

import (
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/entity"
)

type StudentUsecase struct {
	repoStudent internal.StudentRepo
	repoWebinar internal.WebinarRepo
}

func NewStudentUsecase(repoStudent internal.StudentRepo, repoWebinar internal.WebinarRepo) *StudentUsecase {
	return &StudentUsecase{repoStudent: repoStudent, repoWebinar: repoWebinar}
}

func (s *StudentUsecase) StudentSignIn(id int64) error {
	panic("implement me")
}

func (s *StudentUsecase) StudentSignUp(student entity.Student) error {
	panic("implement me")
}

func (s *StudentUsecase) GetStudentProfile(id int64) (entity.Student, error) {
	panic("implement me")
}

func (s *StudentUsecase) SearchWebinarBySlug(slug string) ([]entity.Webinar, error) {
	panic("implement me")
}

func (t *StudentUsecase) GetWebinarByID(id int64) (entity.Webinar, error) {
	panic("implement me")
}

func (s *StudentUsecase) EnrollWebinar(studentID int64, param entity.EnrollWebinarParam) error {
	panic("implement me")
}

func (s *StudentUsecase) GetAllRegisteredWebinar(studentID int64) ([]entity.Webinar, error) {
	panic("implement me")
}

func (s *StudentUsecase) CancelEnrollmentWebinar(studentID int64, param entity.CancelEnrollmentWebinarParam) error {
	panic("implement me")
}

