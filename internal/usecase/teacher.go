package usecase

import (
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/entity"
)

type TeacherUsecase struct {
	repoTeacher internal.TeacherRepo
	repoWebinar internal.WebinarRepo
}

func NewTeacherUsecase(repoTeacher internal.TeacherRepo, repoWebinar internal.WebinarRepo) *TeacherUsecase {
	return &TeacherUsecase{repoTeacher: repoTeacher, repoWebinar: repoWebinar}
}

func (t *TeacherUsecase) TeacherSignIn(id int64) error {
	// TODO: Implement me!
	return entity.ErrUserNotFound
}

func (t *TeacherUsecase) TeacherSignUp(teacher entity.Teacher) error {
	// TODO: Implement me!
	return nil
}

func (t *TeacherUsecase) GetTeacherProfile(id int64) (entity.Teacher, error) {
	// TODO: Implement me!
	return entity.Teacher{Name: "ridwanakf"}, nil
}

func (t *TeacherUsecase) GetAllWebinar(id int64) ([]entity.Webinar, error) {
	// TODO: Implement me!
	return nil, nil
}

func (t *TeacherUsecase) GetWebinarByID(id int64) (entity.Webinar, error) {
	panic("implement me")
}

func (t *TeacherUsecase) CreateNewWebinar(id int64, param entity.CreateWebinarParam) (entity.Webinar, error) {
	panic("implement me")
}

func (t *TeacherUsecase) UpdateWebinar(id int64, param entity.UpdateWebinarParam) error {
	panic("implement me")
}

func (t *TeacherUsecase) DeleteWebinar(id int64, param entity.DeleteWebinarParam) error {
	panic("implement me")
}

func (t *TeacherUsecase) ApproveWaitingList(studentID int64, webinarID int64) error {
	panic("implement me")
}
