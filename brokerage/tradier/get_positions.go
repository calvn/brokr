package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/structs"
	"github.com/calvn/brokr/brokerage/tradier/templates"
	"github.com/calvn/go-tradier/tradier"
)

func (b *Brokerage) GetPositions() (string, error) {
	positions, _, err := b.client.Account.Positions(*b.AccountID)
	if err != nil {
		return "", err
	}

	symbols := make([]string, len(*positions))
	for i, p := range *positions {
		symbols[i] = *p.Symbol
	}

	var quotes *tradier.Quotes
	if len(symbols) > 0 {
		quotes, _, err = b.client.Markets.Quotes(symbols)
		if err != nil {
			return "", err
		}
	}

	pw := structs.NewPositionsWrapper(positions, quotes)

	tmpl := template.Must(template.New("").Funcs(templates.FuncMap()).Parse(templates.PositionsTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, pw)
	output := out.String()

	return output, nil
}
