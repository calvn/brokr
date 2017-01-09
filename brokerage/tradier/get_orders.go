package tradier

import (
	"bytes"
	"html/template"

	"github.com/calvn/go-tradier/tradier"
)

func (b *Brokerage) GetOrders() (string, error) {
	account, _, err := b.client.Account.Orders(*b.AccountID)
	if err != nil {
		return "", err
	}

	var orders *tradier.Orders
	if account.Orders != nil {
		orders = account.Orders
	}

	tmpl := template.Must(template.New("").Parse(ordersTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, orders)
	output := out.String()

	return output, nil
}
