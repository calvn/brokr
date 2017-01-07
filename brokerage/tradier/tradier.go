package tradier

import (
	"fmt"
	"net/http"

	"github.com/calvn/go-tradier/tradier"
)

// Brokerage represents the Tradier brokerage used for the runner
type Brokerage struct {
	client *tradier.Client

	AccountID *string // Account is the account ID that will be used
}

// NewBrokerage creates a new instance of *Brokerage
func NewBrokerage(httpClient *http.Client, accountID string) *Brokerage {
	client := tradier.NewClient(httpClient)

	u, _, err := client.User.Profile()
	if err != nil {
		return nil
	}

	b := &Brokerage{
		client: client,
	}

	// Set sane defaults
	if accountID == "" {
		b.setDefaultAccount(u)
	} else {
		b.AccountID = tradier.String(accountID)
	}

	return b
}

func (b *Brokerage) setDefaultAccount(user *tradier.User) {
	// If no accounts found, set it to a dummy value
	if len(user.Profile.Account) == 0 {
		b.AccountID = tradier.String("UNKNOWN")
	}

	b.AccountID = user.Profile.Account[0].AccountNumber
}

// SwitchAccount switches the account that the client uses to accountID.
func (b *Brokerage) SwitchAccount(accountID string) error {
	p, _, err := b.client.User.Profile()
	if err != nil {
		return err
	}

	for _, account := range *p.Accounts {
		if *account.AccountNumber == accountID {
			*b.AccountID = accountID
			return nil
		}
	}

	return fmt.Errorf("No account %s found for user.", accountID)
}

// Name returns the name of the brokerage.
func (b *Brokerage) Name() string {
	return "Tradier"
}
