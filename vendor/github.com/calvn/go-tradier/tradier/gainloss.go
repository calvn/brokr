package tradier

import (
	"encoding/json"
	"time"
)

type GainLoss struct {
	ClosedPosition []ClosedPosition `json:"closed_position,omitempty"`
}

type gainLoss GainLoss

type ClosedPosition struct {
	ClosedDate      *time.Time `json:"close_date,omitempty"`
	Cost            *float64   `json:"cost,omitempty"`
	GainLoss        *float64   `json:"gain_loss,omitempty"`
	GainLossPercent *float64   `json:"gain_loss_percent,omitempty"`
	OpenDate        *time.Time `json:"open_date,omitempty"`
	Proceeds        *float64   `json:"proceeds,omitempty"`
	Quantity        *int       `json:"quantity,omitempty"`
	Symbol          *string    `json:"symbol,omitempty"`
	Term            *int       `json:"term,omitempty"`
}

type closedPosition struct {
	*ClosedPosition `json:"closed_position,omitempty"`
}

func (g *GainLoss) UnmarshalJSON(b []byte) (err error) {
	gainLossStr := ""
	gainLossObj := gainLoss{}
	closedPositionObj := closedPosition{}

	// If closed_position is a string, i.e. "null"
	if err = json.Unmarshal(b, &gainLossStr); err == nil {
		return nil
	}

	// If closed_position is an array
	if err = json.Unmarshal(b, &gainLossObj); err == nil {
		*g = GainLoss(gainLossObj)
		return nil
	}

	// If closed_position is an object
	if err = json.Unmarshal(b, &closedPositionObj); err == nil {
		*g = GainLoss{
			ClosedPosition: []ClosedPosition{*closedPositionObj.ClosedPosition},
		}
		return nil
	}

	return nil
}

func (g *GainLoss) MarshalJSON() ([]byte, error) {
	// If []ClosedPosition is empty
	if len(g.ClosedPosition) == 0 {
		return json.Marshal("null")
	}

	// If []ClosedPosition is size 1, return first and only object
	if len(g.ClosedPosition) == 1 {
		return json.Marshal(map[string]interface{}{
			"closed_position": g.ClosedPosition[0],
		})
	}

	// Otherwise mashal entire GainLoss object normally
	return json.Marshal(*g)
}
