package tradier

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/calvn/go-tradier/tradier"
)

func (b *Brokerage) GetOrders() error {
	orders, _, err := b.client.Account.Orders(*b.AccountID)
	if err != nil {
		return err
	}

	// FIXME: Correctly print orders
	fmt.Println(orders)

	return nil
}

func (b *Brokerage) PlaceOrder(class, symbol, duration, side string, quantity int, orderType string, price float64) (string, error) {
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

	orders, _, err := b.client.Order.Create(*b.AccountID, params)
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
