package brokerage

import (
	"github.com/calvn/brokr/brokerage/tradier"
	"github.com/calvn/brokr/config"
	"golang.org/x/oauth2"
)

// Broker is the interface for any brokerage
type Broker interface {
	Name() string
	GetQuotes([]string) (string, error)
	GetPositions() error
	GetOrders() error
	PlaceOrder(string, string, string, string, int, string, float64) (string, error) // class, symbol, duration, side, amount, type, limit/stop price
	CancelOrder([]string) error
}

// New creates a new brokerage object based on the provided configuration
func New(config *config.Config) *Broker {
	var b Broker

	switch config.Brokerage {
	case "tradier":
		tokenSource := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: config.AccessToken},
		)

		oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
		b = tradier.NewBrokerage(oauthClient)
	}

	return &b
}
