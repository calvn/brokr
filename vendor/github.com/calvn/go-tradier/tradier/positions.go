package tradier

import (
	"encoding/json"
	"time"
)

type Positions struct {
	Position []Position `json:"position,omitempty"`
}

type positions Positions

type Position struct {
	CostBasis    *float64   `json:"cost_basis,omitempty"`
	DateAcquired *time.Time `json:"date_acquired,omitempty"`
	ID           *int       `json:"id,omitempty"`
	Quantity     *int       `json:"quantity,omitempty"`
	Symbol       *string    `json:"symbol,omitempty"`
}

type position struct {
	*Position `json:"position,omitempty"`
}

func (p *Positions) UnmarshalJSON(b []byte) (err error) {
	positionsStr := ""
	positionsObj := positions{}
	positionObj := position{}

	// If postion is a string, i.e. "null"
	if err = json.Unmarshal(b, &positionsStr); err == nil {
		return nil
	}

	// If position is an array
	if err = json.Unmarshal(b, &positionsObj); err == nil {
		*p = Positions(positionsObj)
		return nil
	}

	// If position is an object
	if err = json.Unmarshal(b, &positionObj); err == nil {
		*p = Positions{
			Position: []Position{*positionObj.Position},
		}
		return nil
	}

	return nil
}

func (p *Positions) MarshalJSON() ([]byte, error) {
	// If []Position is empty
	if len(p.Position) == 0 {
		return json.Marshal("null")
	}

	// If []Position is size 1, return first and only object
	if len(p.Position) == 1 {
		return json.Marshal(map[string]interface{}{
			"position": p.Position[0],
		})
	}

	// Otherwise mashal entire Positions object normally
	return json.Marshal(*p)
}
