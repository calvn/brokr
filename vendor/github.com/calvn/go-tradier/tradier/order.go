package tradier

// OrderService handles routes related to orders
// from the Tradier API.
type OrderService service

// OrderParams specifies the query parameters for querying an order.
// Refer to https://godoc.org/github.com/google/go-querystring/query for building the struct mapping.
type OrderParams struct {
	Class        string  `url:"class"`
	Symbol       string  `url:"symbol"`
	Duration     string  `url:"duration"`
	Side         string  `url:"side,omitempty"`
	Quantity     int     `url:"quantity,omitempty"`
	Type         string  `url:"type"`
	Price        float64 `url:"price,omitempty"`
	Stop         float64 `url:"stop,omitempty"`
	OptionSymbol string  `url:"option_symbol,omitempty"`

	// Specific to multileg orders
	MultiSide         []string `url:"side,omitempty,[]"`
	MultiQuantity     []int    `url:"quantity,omitempty,[]"`
	MultiOptionSymbol []string `url:"option_symbol,omitempty,[]"`

	// Specific to preview
	Preview bool `url:"preview,omitempty"`
}
