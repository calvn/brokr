package tradier

import (
	"bytes"
	"strconv"
	"text/template"

	"github.com/calvn/brokr/brokerage/tradier/structs"
	"github.com/calvn/brokr/brokerage/tradier/templates"
	"github.com/calvn/go-tradier/tradier"
)

// CancelOrder cancels pending orders
func (b *Brokerage) CancelOrder(orderIDs []string) (string, error) {
	// FIXME: Handle errors gracefully. Log on errors and continue on cancelled orders
	orders := &tradier.Orders{}
	for _, id := range orderIDs {
		_order, _, err := b.client.Order.Delete(*b.AccountID, id)
		if err != nil {
			return "", err
		}

		_id := strconv.Itoa(*_order.ID)
		order, _, err := b.client.Account.OrderStatus(*b.AccountID, _id)
		if err != nil {
			return "", err
		}

		*orders = append(*orders, order)
	}
	ow := structs.NewOrdersWrapper(orders)

	tmpl := template.Must(template.New("").Funcs(templates.FuncMap()).Parse(templates.OrdersTemplate))
	var out bytes.Buffer

	tmpl.Execute(&out, ow)
	output := out.String()

	return output, nil
}
