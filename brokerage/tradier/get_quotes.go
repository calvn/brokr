package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/templates"
)

func (b *Brokerage) GetQuotes(symbols []string) (string, error) {
	quotes, _, err := b.client.Markets.Quotes(symbols)
	if err != nil {
		return "", err
	}

	tmpl := template.Must(template.New("").Funcs(templates.FuncMap()).Parse(templates.QuotesTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, quotes)
	output := out.String()

	return output, nil
}
