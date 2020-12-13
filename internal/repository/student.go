package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/weebinar/internal/entity"
)

type StudentDB struct {
	db *sqlx.DB
}

func NewStudentDB(db *sqlx.DB) *StudentDB {
	return &StudentDB{db: db}
}

func (s *StudentDB) IsUserExist(id int64) error {
	isExist := false

	err := s.db.Get(&isExist, SQLCheckStudentExist, id)
	if err != nil || !isExist {
		return entity.ErrUserNotFound
	}

	return nil
}

func (s *StudentDB) InsertNewUser(student entity.Student) error {
	_, err := s.db.Exec(SQLInsertNewStudent, student.ID, student.Name, student.Email, student.Picture)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentDB) GetProfile(id int64) (entity.Student, error) {
	var student entity.Student

	err := s.db.Get(&student, SQLGetStudentProfile, id)
	if err != nil {
		return student, err
	}

	return student, nil
}

func (s *StudentDB) GetAllRegisteredWebinar(id int64) ([]entity.Webinar, error) {
	var webinars []entity.Webinar

	err := s.db.Select(&webinars, SQLGetRegisteredWebinars, id)
	if err != nil {
		return nil, err
	}

	return webinars, nil
}

func (s *StudentDB) EnrollWebinar(id int64, param entity.EnrollWebinarParam) error {
	_, err := s.db.Exec(SQLEnrollWebinar, id, param.WebinarID, param.TeacherID)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentDB) CancelEnrollmentWebinar(id int64, param entity.CancelEnrollmentWebinarParam) error {
	_, err := s.db.Exec(SQLCancelEnrollWebinar, id, param.WebinarID, param.TeacherID)
	if err != nil {
		return err
	}

	return nil
}
