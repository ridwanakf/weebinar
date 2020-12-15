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

func (t *TeacherUsecase) TeacherSignIn(id string) error {
	err := t.repoTeacher.IsUserExist(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TeacherUsecase) TeacherSignUp(teacher entity.Teacher) error {
	err := t.repoTeacher.InsertNewUser(teacher)
	if err != nil {
		return err
	}
	return nil
}

func (t *TeacherUsecase) GetTeacherProfile(id string) (entity.Teacher, error) {
	teacher, err := t.repoTeacher.GetProfile(id)
	if err != nil {
		return entity.Teacher{}, err
	}
	return teacher, nil
}

func (t *TeacherUsecase) GetAllWebinar(teacherID string) ([]entity.Webinar, error) {
	webinars, err := t.repoWebinar.GetAllWebinar(teacherID)
	if err != nil {
		return nil, err
	}

	return webinars, nil
}

func (t *TeacherUsecase) GetWebinarByID(id int64) (entity.Webinar, error) {
	webinar, err := t.repoWebinar.GetWebinarByID(id)
	if err != nil {
		return entity.Webinar{}, err
	}

	return webinar, nil
}

func (t *TeacherUsecase) CreateNewWebinar(id string, param entity.CreateWebinarParam) error {
	err := param.Validate()
	if err != nil {
		return err
	}

	err = t.repoWebinar.InsertNewWebinar(id, param)
	if err != nil {
		return err
	}
	return nil
}

func (t *TeacherUsecase) UpdateWebinar(id string, param entity.UpdateWebinarParam) error {
	err := param.Validate()
	if err != nil {
		return err
	}

	err = t.repoWebinar.UpdateWebinar(id, param)
	if err != nil {
		return err
	}
	return nil
}

func (t *TeacherUsecase) DeleteWebinar(id string, param entity.DeleteWebinarParam) error {
	err := param.Validate()
	if err != nil {
		return err
	}

	err = t.repoWebinar.DeleteWebinar(id, param)
	if err != nil {
		return err
	}
	return nil
}

func (t *TeacherUsecase) ApproveWaitingList(id string, studentID string, webinarID int64) error {
	err := t.repoTeacher.ApproveWaitingList(id, studentID, webinarID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TeacherUsecase) RejectWaitingList(id string, studentID string, webinarID int64) error {
	err := t.repoTeacher.RejectWaitingList(id, studentID, webinarID)
	if err != nil {
		return err
	}

	return nil
}