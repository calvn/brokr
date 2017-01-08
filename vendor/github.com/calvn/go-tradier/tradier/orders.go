package tradier

import (
	"encoding/json"
	"time"
)

// Orders represents the orders JSON object.
type Orders []*Order

// Order represents the `order` JSON object.
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
	PartnerID *string `json:"partner_id,omitempty"`

	// Specific to order preview
	Commission    *float64 `json:"commission,omitempty"`
	Cost          *float64 `json:"cost,omitempty"`
	ExtendedHours *bool    `json:"extended_hours,omitempty"`
	Fees          *float64 `json:"fees,omitempty"`
	MarginChange  *float64 `json:"margin_change,omitempty"`
	Result        *bool    `json:"result,omitempty"`   // Not present in documentation
	Strategy      *string  `json:"strategy,omitempty"` // Not present in documentation, specific to multileg
	unwrapped     bool
}

type order Order

// UnmarshalJSON unmarshals order into Order object.
func (o *Order) UnmarshalJSON(b []byte) error {
	var oc struct {
		*order `json:"order,omitempty"`
	}
	oObj := order{}

	// If wrapped in watchlist object
	if err := json.Unmarshal(b, &oc); err == nil {
		if oc.order != nil {
			*o = Order(*oc.order)
			return nil
		}
	}

	// If not wrapped in anything
	if err := json.Unmarshal(b, &oObj); err == nil {
		*o = Order(oObj)
		return nil
	}

	return nil
}

// MarshalJSON marshals Order into its JSON representation.
func (o *Order) MarshalJSON() ([]byte, error) {
	if o.unwrapped {
		return json.Marshal(*o)
	}

	return json.Marshal(map[string]interface{}{
		"order": *o,
	})
}

// UnmarshalJSON unmarshals orders into Orders object.
func (o *Orders) UnmarshalJSON(b []byte) (err error) {
	var oCollection struct {
		O struct {
			O []*Order `json:"order,omitempty"`
		} `json:"orders,omitempty"`
	}
	var oObject struct {
		O struct {
			O *Order `json:"order,omitempty"`
		} `json:"orders,omitempty"`
	}
	var oNull string

	var oUnwrapped struct {
		O []*Order `json:"order,omitempty"`
	}

	var oSingle struct {
		O *Order `json:"order,omitempty"`
	}

	// If unwrapped from user object
	if err = json.Unmarshal(b, &oUnwrapped); err == nil {
		*o = oUnwrapped.O
		return nil
	}

	// If unwrapped from user object
	if err = json.Unmarshal(b, &oSingle); err == nil {
		*o = Orders{oSingle.O}
		return nil
	}

	// If order is a string, i.e. "null"
	if err = json.Unmarshal(b, &oNull); err == nil {
		return nil
	}

	// If order is a JSON array
	if err = json.Unmarshal(b, &oCollection); err == nil {
		*o = oCollection.O.O
		return nil
	}

	// If order is an object
	if err = json.Unmarshal(b, &oObject); err == nil {
		tmp := make([]*Order, 0)
		tmp = append(tmp, oObject.O.O)
		*o = Orders(tmp)
		return nil
	}

	return nil
}

// MarshalJSON marshals Orders into JSON.
func (o *Orders) MarshalJSON() ([]byte, error) {
	// Set wrapped to true to marshal differently
	for _, order := range *o {
		order.unwrapped = true
	}

	// If []Watchlist is empty
	if len(*o) == 0 {
		return json.Marshal(map[string]interface{}{
			"orders": "null",
		})
	}

	// If []Watchlist is size 1, return first and only object
	if len(*o) == 1 {
		order := *o
		return json.Marshal(map[string]interface{}{
			"order": order[0],
		})
	}

	// Otherwhise marshal normally
	return json.Marshal(map[string]interface{}{
		"order": *o,
	})
}
