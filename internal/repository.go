package internal

import "github.com/ridwanakf/weebinar/internal/entity"

type TeacherRepo interface {
	IsUserExist(id int64) error
	InsertNewUser(teacher entity.Teacher) error
	GetProfile(id int64) (entity.Teacher, error)
	ApproveWaitingList(id int64, studentID int64, webinarID int64) error
	RejectWaitingList(id int64, studentID int64, webinarID int64) error
}

type StudentRepo interface {
	IsUserExist(id int64) error
	InsertNewUser(student entity.Student) error
	GetProfile(id int64) (entity.Student, error)
	GetAllRegisteredWebinar(id int64) ([]entity.Webinar, error)
	EnrollWebinar(id int64, param entity.EnrollWebinarParam) error
	CancelEnrollmentWebinar(id int64, param entity.CancelEnrollmentWebinarParam) error
}

type WebinarRepo interface {
	// Teacher
	GetAllWebinar(id int64) ([]entity.Webinar, error)
	InsertNewWebinar(id int64, param entity.CreateWebinarParam) (entity.Webinar, error)
	UpdateWebinar(id int64, param entity.UpdateWebinarParam) error
	DeleteWebinar(id int64, param entity.DeleteWebinarParam) error

	// Student
	GetWebinarBySlug(slug string) ([]entity.Webinar, error)
	GetWebinarByID(id int64) (entity.Webinar, error)
}
