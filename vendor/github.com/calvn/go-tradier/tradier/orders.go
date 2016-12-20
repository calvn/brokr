package tradier

import (
	"encoding/json"
	"time"
)

// Orders represents the `orders` JSON object
// If `order` is an array of objects, it can be accessed though indexing on Orders.Order
// If `order` is a single object, it can be accessed on the zeroth element
// If `order` is "null", Orders.Order will be empty
type Orders struct {
	Order []Order `json:"order,omitempty"`
}

type orders Orders

// Order represents the `order` JSON object
type Order struct {
	AvgFillPrice      *float64   `json:"avg_fill_price,omitempty"`
	Class             *string    `json:"class,omitempty"`
	CreateDate        *time.Time `json:"create_date,omitempty"`
	Duration          *string    `json:"duration,omitempty"`
	ExecQuantity      *float64   `json:"exec_quantity,omitempty"`
	ID                *int       `json:"id,omitempty"`
	LastFillPrice     *float64   `json:"last_fill_price,omitempty"`
	LastFillQuantity  *float64   `json:"last_fill_quantity,omitempty"`
	Quantity          *float64   `json:"quantity,omitempty"`
	RemainingQuantity *float64   `json:"remaining_quantity,omitempty"`
	Side              *string    `json:"side,omitempty"`
	Status            *string    `json:"status,omitempty"`
	Symbol            *string    `json:"symbol,omitempty"`
	TransactionDate   *time.Time `json:"transaction_date,omitempty"`
	Type              *string    `json:"type,omitempty"`

	// Specific to order creation
	PartnerId *string `json:"partner_id,omitempty"`

	// Specific to order preview
	Commission    *float64 `json:"commission,omitempty"`
	Cost          *float64 `json:"cost,omitempty"`
	ExtendedHours *bool    `json:"extended_hours,omitempty"`
	Fees          *float64 `json:"fees,omitempty"`
	MarginChange  *float64 `json:"margin_change,omitempty"`
	Result        *bool    `json:"result,omitempty"`
}

// NOTE: There should be a better way of handling Order without stubbing
//       Order inside a struct, similar to MarshalJSON via map[string]interface{}
type order struct {
	*Order `json:"order,omitempty"`
}

func (o *Orders) UnmarshalJSON(b []byte) (err error) {
	ordersStr := ""
	ordersObj := orders{}
	orderObj := order{}

	// If order is a string, i.e. "null"
	if err = json.Unmarshal(b, &ordersStr); err == nil {
		return nil
	}

	// If order is a JSON array
	if err = json.Unmarshal(b, &ordersObj); err == nil {
		*o = Orders(ordersObj)
		return nil
	}

	// If order is an object
	if err = json.Unmarshal(b, &orderObj); err == nil {
		*o = Orders{Order: []Order{*orderObj.Order}}
		return nil
	}

	return nil
}

func (o *Orders) MarshalJSON() ([]byte, error) {
	// If []Order is empty
	if len(o.Order) == 0 {
		return json.Marshal("null")
	}

	// If []Order slice is size 1, return first and only object
	if len(o.Order) == 1 {
		return json.Marshal(map[string]interface{}{
			"order": o.Order[0],
		})
	}

	// Otherwise entire mashal Orders object normally
	return json.Marshal(*o)
}
