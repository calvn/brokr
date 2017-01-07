package tradier

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/calvn/go-tradier/tradier"
)

var orderTemplate = `{{range . -}}
{{if .Symbol -}}
Preview order:
{{- else -}}
Order:
{{- end}}
{{- if .ID}}
  OrderID: {{.ID}}
{{- end}}
{{- if .Commission}}
  Commission: {{.Commission}}
{{- end}}
{{- if .Cost}}
  Cost: {{.Cost}}
{{- end}}
{{- if .ExtendedHours}}
  Extended Hours: {{.ExtendedHours}}
{{- end}}
{{- if .Fees}}
  Fees: {{.Fees}}
{{- end}}
{{- if .MarginChange}}
  Margin Change: {{.MarginChange}}
{{- end}}
{{- if .Cost}}
  Cost: {{.Cost}}
{{- end}}
{{- if .Status}}
  Status: {{.Status}}
{{- end}}
{{- end}}`

func (b *Brokerage) GetOrders() error {
	orders, _, err := b.client.Account.Orders(*b.AccountID)
	if err != nil {
		return err
	}

	// FIXME: Correctly print orders
	fmt.Println(orders)

	return nil
}

func (b *Brokerage) CreateOrder(preview bool, class, symbol, duration, side string, quantity int, orderType string, price float64) (string, error) {
	params := &tradier.OrderParams{
		Preview:  preview,
		Class:    class,
		Symbol:   symbol,
		Duration: duration,
		Side:     side,
		Quantity: quantity,
		Type:     orderType,
	}

	switch orderType {
	case "limit":
		params.Price = price
	case "stop":
		params.Stop = price
	}

	order, _, err := b.client.Order.Create(*b.AccountID, params)
	if err != nil {
		return "", err
	}

	tmpl := template.Must(template.New("").Parse(orderTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, order)
	output := out.String()

	return output, nil
}

// FIXME: Order.Delete should return order and not orders
func (b *Brokerage) CancelOrder(orderIDs []string) error {
	// TODO: Implement multi-error, print out order status after cancel submission
	for _, id := range orderIDs {
		_, _, err := b.client.Order.Delete(*b.AccountID, id)
		if err != nil {
			return err
		}
	}

	return nil
}
