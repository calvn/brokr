package brokr

import (
	"github.com/calvn/brokr/brokerage"
	"github.com/calvn/brokr/config"
)

type Runner struct {
	brokerage *brokerage.Brokerage
	config    *config.Config
}

func NewRunner(config *config.Config) *Runner {
	// TODO: Init client based on config
	// Currently defaults to Tradier, the only supported brokerage
	b := brokerage.New(config)

	r := &Runner{
		brokerage: b,
		config:    config,
	}

	return r
}

func (r *Runner) Brokerage() string {
	return (*r.brokerage).Name()
}
