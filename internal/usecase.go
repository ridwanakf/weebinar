package internal

import "github.com/ridwanakf/weebinar/internal/entity"

type TeacherUC interface {
	TeacherSignIn(id string) error
	TeacherSignUp(teacher entity.Teacher) error
	GetTeacherProfile(id string) (entity.Teacher, error)

	GetAllWebinar(teacherID string) ([]entity.Webinar, error)
	GetWebinarByID(id int64) (entity.Webinar, error)
	CreateNewWebinar(id string, param entity.CreateWebinarParam) error
	UpdateWebinar(id string, param entity.UpdateWebinarParam) error
	DeleteWebinar(id string, param entity.DeleteWebinarParam) error

	ApproveWaitingList(id string, studentID string, webinarID int64) error
	RejectWaitingList(id string, studentID string, webinarID int64) error
}

type StudentUC interface {
	StudentSignIn(id string) error
	StudentSignUp(student entity.Student) error
	GetStudentProfile(id string) (entity.Student, error)

	SearchWebinarBySlug(slug string) ([]entity.Webinar, error)
	GetWebinarByID(id int64) (entity.Webinar, error)
	EnrollWebinar(studentID string, param entity.EnrollWebinarParam) error
	GetAllRegisteredWebinar(studentID string) ([]entity.Webinar, error)
	CancelEnrollmentWebinar(studentID string, param entity.CancelEnrollmentWebinarParam) error
}
