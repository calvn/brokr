package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/structs"
	"github.com/calvn/brokr/brokerage/tradier/templates"
)

func (b *Brokerage) GetOrder(id string) (string, error) {
	order, _, err := b.client.Account.OrderStatus(*b.AccountID, id)
	if err != nil {
		return "", err
	}

	// Done to include the amount filled so far
	ot := structs.NewOrderTemplater(order)

	tmpl := template.Must(template.New("").Funcs(templates.FuncMap()).Parse(templates.OrderTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, ot)
	output := out.String()

	return output, nil
}
