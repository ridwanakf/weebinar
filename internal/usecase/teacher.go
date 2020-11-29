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

func (t *TeacherUsecase) TeacherSignIn(id int64) (entity.User, error) {
	panic("implement me")
}

func (t *TeacherUsecase) GetAllWebinar(id int64) ([]entity.Webinar, error) {
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