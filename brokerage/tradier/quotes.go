package tradier

import (
	"fmt"
	"time"
)

func (b *TradierBrokerage) GetQuotes(symbols []string) error {
	quotes, _, err := b.client.Markets.Quotes(symbols)
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
