package internal

import "github.com/ridwanakf/weebinar/internal/entity"

type CommonUC interface {
	CheckUser(id int64) error
	TeacherSignUp() error
	TeacherSignIn() error
	StudentSignUp() error
	StudentSignIn() error
}

type TeacherUC interface {
	SignUp() error
	SignIn() error
	GetAllWebinar(id int64) ([]entity.Webinar, error)
	CreateNewWebinar(id int64, param entity.CreateWebinarParam) (entity.Webinar, error)
	UpdateWebinar(id int64, param entity.UpdateWebinarParam) error
	DeleteWebinar(id int64, param entity.DeleteWebinarParam) error
}

type StudentUC interface {
	SignUp() error
	SignIn() error
	SearchWebinarBySlug(slug string) ([]entity.Webinar, error)
	EnrollWebinar(id int64, param entity.EnrollWebinarParam) error
	GetAllRegisteredWebinar(id int64) ([]entity.StudentWebinar, error)
	CancelEnrollmentWebinar(id int64, param entity.CancelEnrollmentWebinarParam) error
}
