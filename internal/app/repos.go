package app

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ridwanakf/weebinar/internal"
	"github.com/ridwanakf/weebinar/internal/app/config"
	"github.com/ridwanakf/weebinar/internal/repository"
)

type Repos struct {
	TeacherRepo internal.TeacherRepo
	StudentRepo internal.StudentRepo
	WebinarRepo internal.WebinarRepo
}

func newRepos(cfg config.Config) (*Repos, error) {
	db, err := initDB(cfg)
	if err != nil {
		return nil, err
	}

	return &Repos{
		TeacherRepo: repository.NewTeacherDB(db),
		StudentRepo: repository.NewStudentDB(db),
		WebinarRepo: repository.NewWebinarDB(db),
	}, nil
}

func (*Repos) Close() []error {
	var errs []error
	return errs
}

func initDB(cfg config.Config) (*sqlx.DB, error) {
	// Connect MYSQL DB
	db, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.Address)
	if err != nil {
		return nil, err
	}

	// Set db params
	db.SetMaxIdleConns(cfg.Database.MaxConns)
	db.SetMaxOpenConns(cfg.Database.MaxIdleConns)

	return db, nil
}
