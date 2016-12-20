package brokr

import (
	"github.com/calvn/brokr/config"
	"github.com/calvn/go-tradier/tradier"
	"golang.org/x/oauth2"
)

type Runner struct {
	brokerage *Brokerage
	config    *config.Config
}

func NewRunner(config *config.Config) *Runner {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.AccessToken},
	)

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	bClient := tradier.NewClient(oauthClient)

	// TODO: Init client based on config
	// Currently defaults to Tradier, the only supported brokerage
	var brokerage Brokerage
	brokerage = &TradierBrokerage{
		client: bClient,
	}

	r := &Runner{
		brokerage: &brokerage,
		config:    config,
	}

	return r
}
