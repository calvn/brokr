package tradier

import "encoding/json"

// MarketsService handles routes related to orders
// from the Tradier API.
type MarketsService service

// Quotes represents the quotes JSON object.
type Quotes []*Quote

// Quote represents the quote JSON object.
type Quote struct {
	Symbol           *string  `json:"symbol"`
	Description      *string  `json:"description"`
	Exch             *string  `json:"exch"`
	Type             *string  `json:"type"`
	Last             *float64 `json:"last"`
	Change           *float64 `json:"change"`
	ChangePercentage *float64 `json:"change_percentage"`
	Volume           *int     `json:"volume"`
	AverageVolume    *int     `json:"average_volume"`
	LastVolume       *int     `json:"last_volume"`
	TradeDate        *int64   `json:"trade_date"`
	Open             *float64 `json:"open"`
	High             *float64 `json:"high"`
	Low              *float64 `json:"low"`
	Close            *float64 `json:"close"`
	Prevclose        *float64 `json:"prevclose"`
	Week52High       *float64 `json:"week_52_high"`
	Week52Low        *float64 `json:"week_52_low"`
	Bid              *float64 `json:"bid"`
	Bidsize          *int     `json:"bidsize"`
	Bidexch          *string  `json:"bidexch"`
	BidDate          *int64   `json:"bid_date"`
	Ask              *float64 `json:"ask"`
	Asksize          *int     `json:"asksize"`
	Askexch          *string  `json:"askexch"`
	AskDate          *int64   `json:"ask_date"`
	RootSymbols      *string  `json:"root_symbols"`
}

type quote Quote

// UnmarshalJSON unmarshals quote into Quote object.
func (q *Quote) UnmarshalJSON(b []byte) error {
	var qc struct {
		*quote `json:"quote,omitempty"`
	}
	qObj := quote{}

	// If wrapped in quote object
	if err := json.Unmarshal(b, &qc); err == nil {
		if qc.quote != nil {
			*q = Quote(*qc.quote)
			return nil
		}
	}

	// If not wrapped in anything
	if err := json.Unmarshal(b, &qObj); err == nil {
		*q = Quote(qObj)
		return nil
	}

	return nil
}

// MarshalJSON marshals Quote into its JSON representation.
func (q *Quote) MarshalJSON() ([]byte, error) {
	return json.Marshal(*q)
}

// UnmarshalJSON unmarshals quotes into Quotes object.
func (q *Quotes) UnmarshalJSON(b []byte) error {
	var qc struct {
		Q struct {
			Q []*Quote `json:"quote,omitempty"`
		} `json:"quotes,omitempty"`
	}
	var qObj struct {
		Q struct {
			Q *Quote `json:"quote,omitempty"`
		} `json:"quotes,omitempty"`
	}
	var qNull string

	// If quote is null
	if err := json.Unmarshal(b, &qNull); err == nil {
		return nil
	}

	// If quote is a JSON array
	if err := json.Unmarshal(b, &qc); err == nil {
		*q = qc.Q.Q
		return nil
	}
	// If quote is a single object
	if err := json.Unmarshal(b, &qObj); err == nil {
		quotes := Quotes{}
		quotes = append(quotes, qObj.Q.Q)
		*q = quotes
		return nil
	}

	return nil
}

// MarshalJSON marshals Quotes into its JSON representation.
func (q *Quotes) MarshalJSON() ([]byte, error) {
	// If []Quote is empty
	if len(*q) == 0 {
		return json.Marshal(map[string]interface{}{
			"quotes": "null",
		})
	}

	// If []Quote is size 1, return first and only object
	if len(*q) == 1 {
		return json.Marshal(map[string]interface{}{
			"quotes": map[string]interface{}{
				"quote": (*q)[0],
			},
		})
	}

	// Otherwhise marshal normally
	return json.Marshal(map[string]interface{}{
		"quotes": map[string]interface{}{
			"quote": *q,
		},
	})
}
