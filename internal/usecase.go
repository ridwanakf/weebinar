package internal

import "github.com/ridwanakf/weebinar/internal/entity"

type TeacherUC interface {
	TeacherSignIn(id int64) (entity.User, error)
	GetAllWebinar(id int64) ([]entity.Webinar, error)
	CreateNewWebinar(id int64, param entity.CreateWebinarParam) (entity.Webinar, error)
	UpdateWebinar(id int64, param entity.UpdateWebinarParam) error
	DeleteWebinar(id int64, param entity.DeleteWebinarParam) error
}

type StudentUC interface {
	StudentSignIn(id int64) (entity.User, error)
	SearchWebinarBySlug(slug string) ([]entity.Webinar, error)
	EnrollWebinar(id int64, param entity.EnrollWebinarParam) error
	GetAllRegisteredWebinar(id int64) ([]entity.StudentWebinar, error)
	CancelEnrollmentWebinar(id int64, param entity.CancelEnrollmentWebinarParam) error
}
