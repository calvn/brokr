package brokr

import (
	"fmt"

	"github.com/calvn/brokr/config"
	"github.com/calvn/go-tradier/tradier"
	"golang.org/x/oauth2"
)

func GetQuotes(symbols []string) error {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.AccessToken},
	)

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := tradier.NewClient(tc)

	quotes, _, err := client.Quotes.Get(symbols)
	if err != nil {
		return err
	}

	for _, q := range *quotes {
		fmt.Printf("Symbol: %s | Last: %.2f | Low: %.2f | High: %.2f\n", q.Symbol, q.Last, q.Low, q.High)
	}

	return nil
}
