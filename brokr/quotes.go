package brokr

import "fmt"

func (r *Runner) GetQuotes(symbols []string) error {
	quotes, _, err := r.client.Quotes.Get(symbols)
	if err != nil {
		return err
	}

	for _, q := range *quotes {
		fmt.Printf("Symbol: %s | Last: %.2f | Low: %.2f | High: %.2f\n", q.Symbol, q.Last, q.Low, q.High)
	}

	return nil
}
