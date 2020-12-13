package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/weebinar/internal/entity"
)

type WebinarDB struct {
	db *sqlx.DB
}

func NewWebinarDB(db *sqlx.DB) *WebinarDB {
	return &WebinarDB{db: db}
}

func (w *WebinarDB) GetAllWebinar(id int64) ([]entity.Webinar, error) {
	var webinars []entity.Webinar

	err := w.db.Get(&webinars, SQLGetAllWebinarByTeacherID, id)
	if err != nil {
		return nil, err
	}

	return webinars, nil
}

func (w *WebinarDB) InsertNewWebinar(id int64, param entity.CreateWebinarParam) error {
	_, err := w.db.Exec(SQLInsertNewWebinar, id, param.Title, param.Desc, param.Link, param.Category, param.Schedule)
	if err != nil {
		return err
	}

	return nil
}

func (w *WebinarDB) UpdateWebinar(id int64, param entity.UpdateWebinarParam) error {
	_, err := w.db.Exec(SQLUpdateWebinar, param.Title, param.Desc, param.Link, param.Category, param.Schedule, id, param.ID)
	if err != nil {
		return err
	}

	return nil
}

func (w *WebinarDB) DeleteWebinar(id int64, param entity.DeleteWebinarParam) error {
	_, err := w.db.Exec(SQLDeleteWebinar, id, param.ID)
	if err != nil {
		return err
	}

	return nil
}

func (w *WebinarDB) GetWebinarBySlug(slug string) ([]entity.Webinar, error) {
	var webinars []entity.Webinar

	err := w.db.Select(&webinars, SQLGetWebinarBySlug, "%"+slug+"%")
	if err != nil {
		return nil, err
	}

	return webinars, nil
}

func (w *WebinarDB) GetWebinarByID(id int64) (entity.Webinar, error) {
	var webinar entity.Webinar

	err := w.db.Get(&webinar, SQLGetWebinarByID, id)
	if err != nil {
		return webinar, err
	}

	return webinar, nil
}
