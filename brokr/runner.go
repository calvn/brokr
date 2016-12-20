package brokr

import (
	"github.com/calvn/brokr/config"
	"github.com/calvn/go-tradier/tradier"
	"golang.org/x/oauth2"
)

type Runner struct {
	Brokerage *Brokerage
	config    *config.Config
}

func NewRunner(config *config.Config) *Runner {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.AccessToken},
	)

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	tradierClient := tradier.NewClient(oauthClient)

	// brokerage := Brokerage(&TradierBrokerage{
	// 	client: tradierClient,
	// })

	var brokerage Brokerage
	brokerage = &TradierBrokerage{
		client: tradierClient,
	}

	r := &Runner{
		Brokerage: &brokerage,
		config:    config,
	}

	return r
}
