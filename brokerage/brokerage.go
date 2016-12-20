package brokerage

import (
	"github.com/calvn/brokr/brokerage/tradier"
	"github.com/calvn/brokr/config"
	"golang.org/x/oauth2"
)

type Brokerage interface {
	Name() string
	GetQuotes([]string) error
}

func New(config *config.Config) *Brokerage {
	var b Brokerage
	switch config.Brokerage {
	case "tradier":
		tokenSource := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: config.AccessToken},
		)

		oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)

		// TODO: Init client based on config
		// Currently defaults to Tradier, the only supported brokerage
		b = tradier.NewTradierBrokerage(oauthClient)
	}

	return &b
}
