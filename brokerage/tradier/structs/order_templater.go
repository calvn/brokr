package structs

import (
	"github.com/calvn/go-tradier/tradier"
)

// OrderWrapper is wrapper on tradier.Order with the addition of the filled quantity.
type OrderWrapper struct {
	*tradier.Order
	Filled *int
}

// OrdersWrapper is wrapper on tradier.Orders.
type OrdersWrapper []*OrderWrapper

// NewOrderWrapper returns *OrderTemplate with the filled amount calculated.
func NewOrderWrapper(order *tradier.Order) *OrderWrapper {
	ot := &OrderWrapper{Order: order}
	ot.Filled = tradier.Int(int(*order.Quantity - *order.RemainingQuantity))
	// If order has not started to before cancelling, return 0 as filled
	if *order.Status == "canceled" && *order.RemainingQuantity == 0 {
		ot.Filled = tradier.Int(0)
	}
	return ot
}

// NewOrdersWrapper returns *OrdersTemplate with the filled
// amount calculated for earch *OrderTemplate.
func NewOrdersWrapper(orders *tradier.Orders) *OrdersWrapper {
	ot := make(OrdersWrapper, len(*orders))
	for i, o := range *orders {
		ot[i] = NewOrderWrapper(o)
	}

	return &ot
}
