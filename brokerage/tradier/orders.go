package tradier

import (
	"strconv"

	"github.com/calvn/go-tradier/tradier"
)

func (b *TradierBrokerage) PlaceOrder(class, symbol, duration, side string, quantity int, orderType string, price float64) ([]string, error) {
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
		return nil, err
	}

	orderIDs := []string{}

	for _, o := range *orders {
		orderIDs = append(orderIDs, strconv.Itoa(*o.ID))
	}

	return orderIDs, nil
}
