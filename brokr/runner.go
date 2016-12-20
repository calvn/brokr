package brokr

import (
	"github.com/calvn/brokr/brokr/brokerage"
	"github.com/calvn/brokr/config"
	"golang.org/x/oauth2"
)

type Runner struct {
	brokerage *brokerage.Brokerage
	config    *config.Config
}

func NewRunner(config *config.Config) *Runner {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.AccessToken},
	)

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)

	// TODO: Init client based on config
	// Currently defaults to Tradier, the only supported brokerage
	var b brokerage.Brokerage
	b = brokerage.NewTradierBrokerage(oauthClient)

	r := &Runner{
		brokerage: &b,
		config:    config,
	}

	return r
}

func (r *Runner) Brokerage() string {
	return (*r.brokerage).Name()
}
