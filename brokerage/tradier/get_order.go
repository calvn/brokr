package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/templates"
)

func (b *Brokerage) GetOrder(id string) (string, error) {
	order, _, err := b.client.Account.OrderStatus(*b.AccountID, id)
	if err != nil {
		return "", err
	}

	tmpl := template.Must(template.New("").Funcs(templates.FuncMap()).Parse(templates.OrderTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, order)
	output := out.String()

	return output, nil
}
