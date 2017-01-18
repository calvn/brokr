package structs

import (
	"github.com/calvn/go-tradier/tradier"
)

// OrderWrapper is the wrapper on tradier.Order with the addition of the filled quantity.
type OrderWrapper struct {
	*tradier.Order
	Filled *int
}

// OrdersWrapper holds a collection of OrderWrapper.
type OrdersWrapper []*OrderWrapper

// NewOrderWrapper returns *OrderTemplate with the filled amount calculated.
func NewOrderWrapper(order *tradier.Order) *OrderWrapper {
	o := &OrderWrapper{Order: order}
	o.Filled = tradier.Int(int(*order.Quantity - *order.RemainingQuantity))
	// If order has not started to before cancelling, return 0 as filled
	if *order.Status == "canceled" && *order.RemainingQuantity == 0 {
		o.Filled = tradier.Int(0)
	}
	return o
}

// NewOrdersWrapper returns *OrdersTemplate with the filled
// amount calculated for earch *OrderTemplate.
func NewOrdersWrapper(orders *tradier.Orders) *OrdersWrapper {
	o := make(OrdersWrapper, len(*orders))
	for i, order := range *orders {
		o[i] = NewOrderWrapper(order)
	}

	return &o
}
