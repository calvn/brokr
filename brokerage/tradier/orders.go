package tradier

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/calvn/go-tradier/tradier"
)

func (b *TradierBrokerage) GetOrders() error {
	orders, _, err := b.client.Account.Orders(*b.Account)
	if err != nil {
		return err
	}

	// FIXME: Correctly print orders
	fmt.Println(orders)

	return nil
}

func (b *TradierBrokerage) PlaceOrder(class, symbol, duration, side string, quantity int, orderType string, price float64) (string, error) {
	params := &tradier.OrderParams{
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

	orders, _, err := b.client.Order.Create(*b.Account, params)
	if err != nil {
		return "", err
	}

	orderIDs := []string{}

	for _, o := range *orders {
		orderIDs = append(orderIDs, strconv.Itoa(*o.ID))
	}

	output := strings.Join(orderIDs, "\n")

	return output, nil
}

// FIXME: Order.Delete should return order and not orders
func (b *TradierBrokerage) CancelOrder(orderIDs []string) error {
	// TODO: Implement multi-error, print out order status after cancel submission
	for _, id := range orderIDs {
		_, _, err := b.client.Order.Delete(*b.Account, id)
		if err != nil {
			return err
		}
	}

	return nil
}
