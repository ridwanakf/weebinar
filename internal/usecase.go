package internal

import "github.com/ridwanakf/weebinar/internal/entity"

type TeacherUC interface {
	TeacherSignIn(id int64) error
	TeacherSignUp(teacher entity.Teacher) error
	GetTeacherProfile(id int64) (entity.Teacher, error)

	GetAllWebinar(teacherID int64) ([]entity.Webinar, error)
	GetWebinarByID(id int64) (entity.Webinar, error)
	CreateNewWebinar(id int64, param entity.CreateWebinarParam) (entity.Webinar, error)
	UpdateWebinar(id int64, param entity.UpdateWebinarParam) error
	DeleteWebinar(id int64, param entity.DeleteWebinarParam) error

	ApproveWaitingList(studentID int64, webinarID int64) error
}

type StudentUC interface {
	StudentSignIn(id int64) error
	StudentSignUp(student entity.Student) error
	GetStudentProfile(id int64) (entity.Student, error)

	SearchWebinarBySlug(slug string) ([]entity.Webinar, error)
	GetWebinarByID(id int64) (entity.Webinar, error)
	EnrollWebinar(studentID int64, param entity.EnrollWebinarParam) error
	GetAllRegisteredWebinar(studentID int64) ([]entity.Webinar, error)
	CancelEnrollmentWebinar(studentID int64, param entity.CancelEnrollmentWebinarParam) error
}
