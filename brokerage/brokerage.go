package brokerage

import (
	"github.com/calvn/brokr/brokerage/tradier"
	"github.com/calvn/brokr/config"
)

// Broker is the interface for any brokerage
type Broker interface {
	GetQuotes([]string) (string, error)
	GetPositions() (string, error)
	GetOrders() (string, error)

	// CreateOrder accepts: isPreview, class, symbol, duration, side, amount, type, limit/stop price
	CreateOrder(bool, string, string, string, string, int, string, float64) (string, error)
	GetOrder(string) (string, error)
	CancelOrder([]string) (string, error)
}

// New creates a new brokerage object based on the provided configuration
func New(config *config.Config) *Broker {
	var b Broker

	switch {
	case config.Brokerage == "tradier" && config.Tradier != nil:
		b = tradier.NewBrokerage(config.Tradier)
	}

	return &b
}
