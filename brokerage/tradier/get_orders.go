package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/structs"
	"github.com/calvn/brokr/brokerage/tradier/templates"
)

func (b *Brokerage) GetOrders() (string, error) {
	orders, _, err := b.client.Account.Orders(*b.AccountID)
	if err != nil {
		return "", err
	}

	// Done to include the amount filled so far
	ot := structs.NewOrdersTemplater(orders)

	tmpl := template.Must(template.New("").Funcs(templates.FuncMap()).Parse(templates.OrdersTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, ot)
	output := out.String()

	return output, nil
}
