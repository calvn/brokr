package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/templates"
	"github.com/calvn/go-tradier/tradier"
)

func (b *Brokerage) CreateOrder(preview bool, class, symbol, duration, side string, quantity int, orderType string, price float64) (string, error) {
	params := &tradier.OrderParams{
		Preview:  preview,
		Class:    class,
		Symbol:   symbol,
		Duration: duration,
		Side:     side,
		Quantity: quantity,
		Type:     orderType,
	}

	switch orderType {
	case "limit":
		params.Price = price
	case "stop":
		params.Stop = price
	}

	order, _, err := b.client.Order.Create(*b.AccountID, params)
	if err != nil {
		return "", err
	}

	tmpl := template.Must(template.New("").Parse(templates.OrderTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, order)
	output := out.String()

	return output, nil
}
