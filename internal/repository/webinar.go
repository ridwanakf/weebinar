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
	panic("implement me")
}

func (w *WebinarDB) InsertNewWebinar(id int64, param entity.CreateWebinarParam) (entity.Webinar, error) {
	panic("implement me")
}

func (w *WebinarDB) UpdateWebinar(id int64, param entity.UpdateWebinarParam) error {
	panic("implement me")
}

func (w *WebinarDB) DeleteWebinar(id int64, param entity.DeleteWebinarParam) error {
	panic("implement me")
}

func (w *WebinarDB) GetWebinarBySlug(slug string) ([]entity.Webinar, error) {
	panic("implement me")
}

func (w *WebinarDB) GetWebinarByID(id int64) (entity.Webinar, error) {
	panic("implement me")
}
