package tradier

import (
	"bytes"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/templates"
)

func (b *Brokerage) GetPositions() (string, error) {
	positions, _, err := b.client.Account.Positions(*b.AccountID)
	if err != nil {
		return "", err
	}

	tmpl := template.Must(template.New("").Parse(templates.PositionsTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, positions)
	output := out.String()

	return output, nil
}
