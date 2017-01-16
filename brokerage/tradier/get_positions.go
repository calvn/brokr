package tradier

import (
	"bytes"
	"log"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/templates"
)

func (b *Brokerage) GetPositions() (string, error) {
	positions, _, err := b.client.Account.Positions(*b.AccountID)
	if err != nil {
		return "", err
	}

	log.Println(*(*positions)[0].CostBasis)

	tmpl := template.Must(template.New("").Funcs(templates.FuncMap()).Parse(templates.PositionsTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, positions)
	output := out.String()

	return output, nil
}
