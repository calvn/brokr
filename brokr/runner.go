package brokr

import (
	"github.com/calvn/brokr/config"
	"github.com/calvn/go-tradier/tradier"
	"golang.org/x/oauth2"
)

type Runner struct {
	client *tradier.Client
	config *config.Config
}

func NewRunner(config *config.Config) *Runner {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.AccessToken},
	)

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	tradierClient := tradier.NewClient(oauthClient)

	r := &Runner{
		client: tradierClient,
		config: config,
	}

	return r
}
