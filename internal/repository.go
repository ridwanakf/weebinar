package internal

import "github.com/ridwanakf/weebinar/internal/entity"

type TeacherRepo interface {
	IsUserExist(id string) error
	InsertNewUser(teacher entity.Teacher) error
	GetProfile(id string) (entity.Teacher, error)
	ApproveWaitingList(id string, studentID string, webinarID int64) error
	RejectWaitingList(id string, studentID string, webinarID int64) error
}

type StudentRepo interface {
	IsUserExist(id string) error
	InsertNewUser(student entity.Student) error
	GetProfile(id string) (entity.Student, error)
	GetAllRegisteredWebinar(id string) ([]entity.Webinar, error)
	EnrollWebinar(id string, param entity.EnrollWebinarParam) error
	CancelEnrollmentWebinar(id string, param entity.CancelEnrollmentWebinarParam) error
}

type WebinarRepo interface {
	// Teacher
	GetAllWebinar(id string) ([]entity.Webinar, error)
	InsertNewWebinar(id string, param entity.CreateWebinarParam) error
	UpdateWebinar(id string, param entity.UpdateWebinarParam) error
	DeleteWebinar(id string, param entity.DeleteWebinarParam) error

	// Student
	GetWebinarBySlug(slug string) ([]entity.Webinar, error)
	GetWebinarByID(id int64) (entity.Webinar, error)
}
