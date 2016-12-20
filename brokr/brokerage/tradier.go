package brokerage

import (
	"fmt"
	"net/http"
	"time"

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

func (b *TradierBrokerage) GetQuotes(symbols []string) error {
	quotes, _, err := b.client.Quotes.Get(symbols)
	if err != nil {
		return err
	}

	for _, q := range *quotes {
		time := time.Unix(0, *q.TradeDate*int64(time.Millisecond))
		fmt.Printf(`Symbol: %s
    Last:      %.2f
    PrevClose: %.2f
    Change:    %.2f (%.2f%%)
    Low:       %.2f
    High:      %.2f
    Updated:   %s`, *q.Symbol, *q.Last, *q.Prevclose, *q.Change, *q.ChangePercentage, *q.Low, *q.High, time)
	}

	return nil
}

func (b *TradierBrokerage) Name() string {
	return "Tradier"
}
