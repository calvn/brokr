package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/templates"
)

func (b *Brokerage) GetOrders() (string, error) {
	orders, _, err := b.client.Account.Orders(*b.AccountID)
	if err != nil {
		return "", err
	}

	tmpl := template.Must(template.New("").Parse(templates.OrdersTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, orders)
	output := out.String()

	return output, nil
}
