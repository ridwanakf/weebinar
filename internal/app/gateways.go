package app

import (
	"github.com/ridwanakf/weebinar/internal/app/config"
)

type Gateways struct {
}

func newGateways(cfg *config.Config) *Gateways {
	return &Gateways{
	}
}

func (*Gateways) Close() []error {
	var errs []error
	return errs
}
