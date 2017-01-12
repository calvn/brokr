package tradier

import (
	"encoding/json"
	"log"
	"time"
)

// Positions represents the positions JSON object.
type Positions []*Position

type positions Positions

// Position represents the position JSON object.
type Position struct {
	CostBasis    *float64   `json:"cost_basis,omitempty"`
	DateAcquired *time.Time `json:"date_acquired,omitempty"`
	ID           *int       `json:"id,omitempty"`
	Quantity     *float64   `json:"quantity,omitempty"`
	Symbol       *string    `json:"symbol,omitempty"`
}

// UnmarshalJSON unmarshals positions into Positions object.
func (p *Positions) UnmarshalJSON(b []byte) error {
	var posCol struct {
		P struct {
			P []*Position `json:"position,omitempty"`
		} `json:"positions,omitempty"`
	}
	var posObj struct {
		P struct {
			P *Position `json:"position,omitempty"`
		} `json:"positions,omitempty"`
	}
	var posStr struct {
		P string `json:"positions,omitempty"`
	}
	var err error

	// If positions is null
	if err = json.Unmarshal(b, &posStr); err == nil {
		return nil
	}

	// If watchlist is a JSON array
	if err = json.Unmarshal(b, &posCol); err == nil {
		*p = posCol.P.P
		return nil
	}

	// If watchlist is a single object
	if err = json.Unmarshal(b, &posObj); err == nil {
		pos := make([]*Position, 1)
		pos[0] = posObj.P.P
		*p = Positions(pos)
		return nil
	}

	return err
}

// MarshalJSON marshals Positions into JSON.
func (p *Positions) MarshalJSON() ([]byte, error) {
	log.Printf("%+v\n", p)
	// If []Position is empty
	if len(*p) == 0 {
		return json.Marshal(map[string]interface{}{
			"positions": "null",
		})
	}

	// If []Position is size 1, return first and only object
	if len(*p) == 1 {
		return json.Marshal(map[string]interface{}{
			"positions": map[string]interface{}{
				"position": *(*p)[0],
			},
		})
	}

	// Otherwise wrap and mashal normally
	return json.Marshal(map[string]interface{}{
		"positions": map[string]interface{}{
			"position": *p,
		},
	})
}
