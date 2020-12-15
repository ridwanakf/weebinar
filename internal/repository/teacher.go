package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/weebinar/internal/entity"
)

type TeacherDB struct {
	db *sqlx.DB
}

func NewTeacherDB(db *sqlx.DB) *TeacherDB {
	return &TeacherDB{db: db}
}

func (t *TeacherDB) IsUserExist(id string) error {
	isExist := false

	err := t.db.Get(&isExist, SQLCheckTeacherExist, id)
	if err != nil || !isExist {
		return entity.ErrUserNotFound
	}

	return nil
}

func (t *TeacherDB) InsertNewUser(teacher entity.Teacher) error {
	_, err := t.db.Exec(SQLInsertNewTeacher, teacher.ID, teacher.Name, teacher.Email, teacher.Picture)
	if err != nil {
		return err
	}

	return nil
}

func (t *TeacherDB) GetProfile(id string) (entity.Teacher, error) {
	var teacher entity.Teacher

	err := t.db.Get(&teacher, SQLGetTeacherProfile, id)
	if err != nil {
		return teacher, err
	}

	return teacher, nil
}

func (t *TeacherDB) ApproveWaitingList(id string, studentID string, webinarID int64) error {
	_, err := t.db.Exec(SQLApproveStudent, id, webinarID, studentID)
	if err != nil {
		return err
	}

	return nil
}

func (t *TeacherDB) RejectWaitingList(id string, studentID string, webinarID int64) error {
	_, err := t.db.Exec(SQLRejectStudent, id, webinarID, studentID)
	if err != nil {
		return err
	}

	return nil
}