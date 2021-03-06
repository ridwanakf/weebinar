package app

import (
	"github.com/ridwanakf/weebinar/constant"
	"github.com/ridwanakf/weebinar/internal/app/config"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type WeebinarApp struct {
	Repos    *Repos
	UseCases *Usecases

	Cfg config.Config
}

func NewWeebinarApp() (*WeebinarApp, error) {
	cfg, err := readConfig(constant.ConfigProjectFilepath)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting config")
	}

	app := new(WeebinarApp)

	app.Cfg = cfg

	app.Repos, err = newRepos(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "error initializing repo")
	}

	app.UseCases = newUsecases(app.Repos, cfg)

	return app, nil
}

func (a *WeebinarApp) Close() []error {
	var errs []error

	errs = append(errs, a.Repos.Close()...)
	errs = append(errs, a.UseCases.Close()...)

	return errs
}

func readConfig(cfgPath string) (config.Config, error) {
	f, err := os.Open(cfgPath)
	if err != nil {
		return config.Config{}, errors.Wrapf(err, "config file not found")
	}
	defer f.Close()

	var cfg config.Config

	// Read from config file
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return config.Config{}, errors.Wrapf(err, "error reading config from file")
	}

	// Replace vars that exist in ENV
	if err := env.Parse(&cfg); err != nil {
		return config.Config{}, errors.Wrapf(err, "error reading config from ENV")
	}

	return cfg, nil
}
