package tradier

import (
	"encoding/json"
	"time"
)

type Accounts struct {
	Account []Account `json:"account,omitempty"`
}

type accounts Accounts

type Account struct {
	AccountNumber *string `json:"account_number,omitempty"`

	// Specific to orders
	Orders *Orders `json:"orders,omitempty"`

	// Specific to profile
	Classification *string    `json:"classification,omitempty"`
	DateCreated    *time.Time `json:"date_created,omitempty"`
	DayTrader      *bool      `json:"day_trader,omitempty"`
	OptionLevel    *int       `json:"option_level,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Type           *string    `json:"type,omitempty"`
	LastUpateDate  *time.Time `json:"last_update_date,omitempty"`

	// Specific to positions
	Positions *Positions `json:"positions,omitempty"`

	// Specific to gainloss
	GainLoss *GainLoss `json:"gainloss,omitempty"`

	// Specific to history
	History *History `json:"history,omitempty"`

	// Specific to balances
	Balances *Balances `json:"balances,omitempty"`
}

type account struct {
	*Account `json:"account,omitempty"`
}

func (a *Accounts) UnmarshalJSON(b []byte) (err error) {
	accountsStr := ""
	accountsObj := accounts{}
	accountObj := account{}

	// If account is a string, i.e. "null"
	if err = json.Unmarshal(b, &accountsStr); err == nil {
		return nil
	}

	// If account is an array
	if err = json.Unmarshal(b, &accountsObj); err == nil {
		*a = Accounts(accountsObj)
		return nil
	}

	// If account is an object
	if err = json.Unmarshal(b, &accountObj); err == nil {
		*a = Accounts{
			Account: []Account{*accountObj.Account},
		}
		return nil
	}

	return nil
}

func (a *Accounts) MarshalJSON() ([]byte, error) {
	// If Account is null
	if len(a.Account) == 0 {
		return json.Marshal("null")
	}

	// If Account slice is size 1, return object directly
	if len(a.Account) == 1 {
		return json.Marshal(map[string]interface{}{
			"account": a.Account[0],
		})
	}

	// Otherwise mashal Account normally, in this case using []Order
	return json.Marshal(*a)
}
