package brokr

import (
	"github.com/calvn/brokr/brokerage"
	"github.com/calvn/brokr/config"
)

type Runner struct {
	brokerage *brokerage.Brokerage
	config    *config.Config
}

// NewRunner create a new instance of Runner with the provided configuration
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

// Brokerage returns the brokerage name that is being used by the runner
func (r *Runner) Brokerage() string {
	return (*r.brokerage).Name()
}

// Config returns the runner's config.Config
func (r *Runner) Config() config.Config {
	return *r.config
}
