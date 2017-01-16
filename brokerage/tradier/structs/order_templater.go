package structs

import (
	"github.com/calvn/go-tradier/tradier"
)

// OrderTemplater is wrapper on tradier.Order with the addition of the filled quantity.
type OrderTemplater struct {
	*tradier.Order
	Filled *int
}

// OrdersTemplater is wrapper on tradier.Orders.
type OrdersTemplater []*OrderTemplater

// NewOrderTemplater returns *OrderTemplate with the filled amount calculated.
func NewOrderTemplater(order *tradier.Order) *OrderTemplater {
	ot := &OrderTemplater{Order: order}
	ot.Filled = tradier.Int(int(*order.Quantity - *order.RemainingQuantity))
	// If order has not started to before cancelling, return 0 as filled
	if *order.Status == "canceled" && *order.RemainingQuantity == 0 {
		ot.Filled = tradier.Int(0)
	}
	return ot
}

// NewOrdersTemplater returns *OrdersTemplate with the filled
// amount calculated for earch *OrderTemplate.
func NewOrdersTemplater(orders *tradier.Orders) *OrdersTemplater {
	ot := make(OrdersTemplater, len(*orders))
	for i, o := range *orders {
		ot[i] = NewOrderTemplater(o)
	}

	return &ot
}
