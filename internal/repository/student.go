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
	panic("implement me")
}

func (s *StudentDB) InsertNewUser(student entity.Student) error {
	panic("implement me")
}

func (s *StudentDB) GetProfile(id int64) (entity.Student, error) {
	panic("implement me")
}

func (s *StudentDB) GetAllRegisteredWebinar(id int64) ([]entity.Webinar, error) {
	panic("implement me")
}

func (s *StudentDB) EnrollWebinar(id int64, param entity.EnrollWebinarParam) error {
	panic("implement me")
}

func (s *StudentDB) CancelEnrollmentWebinar(id int64, param entity.CancelEnrollmentWebinarParam) error {
	panic("implement me")
}
