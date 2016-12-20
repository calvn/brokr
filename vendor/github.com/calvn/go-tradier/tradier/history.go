package tradier

import (
	"encoding/json"
	"time"
)

type History struct {
	Event []Event `json:"event,omitempty"`
}

type history History

type Event struct {
	Amount   *float64   `json:"amount,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
	Interest *struct {
		Description *string `json:"description,omitempty"`
		Quantity    *int    `json:"quantity,omitempty"`
	} `json:"interest,omitempty"`
	Journal *struct {
		Quantity *int `json:"quantity,omitempty"`
	} `json:"journal,omitempty"`
	Option *struct {
		Description *string `json:"description,omitempty"`
		OptionType  *string `json:"option_type,omitempty"`
		Quantity    *int    `json:"quantity,omitempty"`
	} `json:"option,omitempty"`
	Type  *string `json:"type,omitempty"`
	Trade *struct {
		Commission  *float64 `json:"commission,omitempty"`
		Description *string  `json:"description,omitempty"`
		Price       *float64 `json:"price,omitempty"`
		Quantity    *int     `json:"quantity,omitempty"`
		Symbol      *string  `json:"symbol,omitempty"`
		TradeType   *string  `json:"trade_type,omitempty"`
	} `json:"trade,omitempty"`
}

type event struct {
	*Event `json:"event,omitempty"`
}

func (h *History) UnmarshalJSON(b []byte) (err error) {
	historyStr := ""
	historyObj := history{}
	eventObj := event{}

	// If event is a string, i.e. "null"
	if err = json.Unmarshal(b, &historyStr); err == nil {
		return nil
	}

	// If event is an array
	if err = json.Unmarshal(b, &historyObj); err == nil {
		*h = History(historyObj)
		return nil
	}

	// If event is an object
	if err = json.Unmarshal(b, &eventObj); err == nil {
		*h = History{
			Event: []Event{*eventObj.Event},
		}
		return nil
	}

	return nil
}

func (h *History) MarshalJSON() ([]byte, error) {
	// If []Event is empty
	if len(h.Event) == 0 {
		return json.Marshal("null")
	}

	// If []Event is size 1, return first and only object
	if len(h.Event) == 1 {
		return json.Marshal(map[string]interface{}{
			"event": h.Event[0],
		})
	}

	// Otherwise mashal entire History object normally
	return json.Marshal(*h)
}
