package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/templates"
)

// CancelOrder cancels pending orders
func (b *Brokerage) CancelOrder(orderIDs []string) (string, error) {
	output := ""

	// FIXME: Append error to output
	for _, id := range orderIDs {
		order, _, err := b.client.Order.Delete(*b.AccountID, id)
		if err != nil {
			return "", err
		}

		var out bytes.Buffer
		tmpl := template.Must(template.New("").Parse(templates.OrderTemplate))
		tmpl.Execute(&out, order)

		output += out.String()
	}

	return output, nil
}
