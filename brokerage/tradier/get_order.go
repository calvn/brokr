package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/structs"
	"github.com/calvn/brokr/brokerage/tradier/templates"
)

// GetOrder fetches a specific order and returns information about it.
func (b *Brokerage) GetOrder(id string) (string, error) {
	order, _, err := b.client.Account.OrderStatus(*b.AccountID, id)
	if err != nil {
		return "", err
	}

	// Done to include the amount filled so far
	ow := structs.NewOrderWrapper(order)

	tmpl := template.Must(template.New("").Funcs(templates.FuncMap()).Parse(templates.OrderTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, ow)
	output := out.String()

	return output, nil
}
