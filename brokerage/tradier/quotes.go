package tradier

import (
	"fmt"
	"strings"
	"time"
)

func (b *TradierBrokerage) GetQuotes(symbols []string) (string, error) {
	quotes, _, err := b.client.Markets.Quotes(symbols)
	if err != nil {
		return "", err
	}

	// Native concat (+=) is more performant than strings.Join but this is done for the newline, for now.
	// http://herman.asia/efficient-string-concatenation-in-go
	entries := []string{}
	for _, q := range *quotes {
		time := time.Unix(0, *q.TradeDate*int64(time.Millisecond))
		entry := fmt.Sprintf(`Symbol: %s
    Last:      %.2f
    PrevClose: %.2f
    Change:    %.2f (%.2f%%)
    Low:       %.2f
    High:      %.2f
    Updated:   %s`, *q.Symbol, *q.Last, *q.Prevclose, *q.Change, *q.ChangePercentage, *q.Low, *q.High, time)

		entries = append(entries, entry)
	}

	output := strings.Join(entries, "\n\n")

	return output, nil
}
