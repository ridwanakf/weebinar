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

func (s *StudentDB) IsUserExist(id string) error {
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

func (s *StudentDB) GetProfile(id string) (entity.Student, error) {
	var student entity.Student

	err := s.db.Get(&student, SQLGetStudentProfile, id)
	if err != nil {
		return student, err
	}

	return student, nil
}

func (s *StudentDB) GetAllRegisteredWebinar(id string) ([]entity.Webinar, error) {
	var (
		webinars      []entity.Webinar
		participants []entity.Participants
	)

	err := s.db.Select(&webinars, SQLGetRegisteredWebinars, id)
	if err != nil {
		return nil, err
	}

	for k, _ := range webinars{
		err = s.db.Select(&participants, SQLGetParticipants, webinars[k].ID)
		if err != nil {
			return webinars, err
		}
		webinars[k].Participants = participants
	}

	return webinars, nil
}

func (s *StudentDB) EnrollWebinar(id string, param entity.EnrollWebinarParam) error {
	_, err := s.db.Exec(SQLEnrollWebinar, id, param.WebinarID, param.TeacherID)
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentDB) CancelEnrollmentWebinar(id string, param entity.CancelEnrollmentWebinarParam) error {
	_, err := s.db.Exec(SQLCancelEnrollWebinar, id, param.WebinarID, param.TeacherID)
	if err != nil {
		return err
	}

	return nil
}
