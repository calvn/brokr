package brokr

import (
	"github.com/calvn/brokr/brokerage"
	"github.com/calvn/brokr/config"
)

type Runner struct {
	brokerage.Brokerage
	config *config.Config
}

// NewRunner create a new instance of Runner with the provided configuration
func NewRunner(config *config.Config) *Runner {
	// TODO: Init client based on config
	// Currently defaults to Tradier, the only supported brokerage
	b := brokerage.New(config)

	r := &Runner{*b, config}

	return r
}

// Config returns the runner's config.Config
func (r *Runner) Config() config.Config {
	return *r.config
}
