package tradier

import (
	"net/http"

	"github.com/calvn/go-tradier/tradier"
)

type TradierBrokerage struct {
	client *tradier.Client
}

func NewTradierBrokerage(httpClient *http.Client) *TradierBrokerage {
	client := tradier.NewClient(httpClient)

	return &TradierBrokerage{
		client: client,
	}
}

func (b *TradierBrokerage) Name() string {
	return "Tradier"
}
