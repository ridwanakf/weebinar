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
	err := s.repoStudent.IsUserExist(id)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *StudentUsecase) StudentSignUp(student entity.Student) error {
	err := s.repoStudent.InsertNewUser(student)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentUsecase) GetStudentProfile(id int64) (entity.Student, error) {
	student, err := s.repoStudent.GetProfile(id)
	if err != nil {
		return entity.Student{}, err
	}

	return student, nil
}

func (s *StudentUsecase) SearchWebinarBySlug(slug string) ([]entity.Webinar, error) {
	webinars, err := s.repoWebinar.GetWebinarBySlug(slug)
	if err != nil {
		return nil, err
	}

	return webinars, nil
}

func (s *StudentUsecase) GetWebinarByID(id int64) (entity.Webinar, error) {
	webinar, err := s.repoWebinar.GetWebinarByID(id)
	if err != nil {
		return entity.Webinar{}, err
	}

	return webinar, nil
}

func (s *StudentUsecase) EnrollWebinar(studentID int64, param entity.EnrollWebinarParam) error {
	err := param.Validate()
	if err != nil {
		return err
	}

	err = s.repoStudent.EnrollWebinar(studentID, param)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentUsecase) GetAllRegisteredWebinar(studentID int64) ([]entity.Webinar, error) {
	webinars, err := s.repoStudent.GetAllRegisteredWebinar(studentID)
	if err != nil {
		return nil, err
	}

	return webinars, nil
}

func (s *StudentUsecase) CancelEnrollmentWebinar(studentID int64, param entity.CancelEnrollmentWebinarParam) error {
	err := param.Validate()
	if err != nil {
		return err
	}

	err = s.repoStudent.CancelEnrollmentWebinar(studentID, param)
	if err != nil {
		return err
	}

	return nil
}
