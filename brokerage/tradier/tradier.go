package tradier

import (
	"fmt"

	"golang.org/x/oauth2"

	"github.com/calvn/brokr/config"
	"github.com/calvn/go-tradier/tradier"
)

// Brokerage represents the Tradier brokerage used for the runner
type Brokerage struct {
	client *tradier.Client

	AccountID *string // Account is the account ID that will be used
}

// NewBrokerage creates a new instance of *Brokerage
func NewBrokerage(config *config.TradierConfig) *Brokerage {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.AccessToken},
	)

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)

	client := tradier.NewClient(oauthClient)

	b := &Brokerage{
		client: client,
	}

	// Set sane defaults
	if config.AccountID == "" {
		b.setDefaultAccount()
	} else {
		b.AccountID = tradier.String(config.AccountID)
	}

	return b
}

func (b *Brokerage) setDefaultAccount() {
	// If no accounts found, set it to a dummy value
	u, _, err := b.client.User.Profile()
	if err != nil || len(u.Profile.Account) == 0 {
		b.AccountID = tradier.String("UNKNOWN")
	}

	b.AccountID = u.Profile.Account[0].AccountNumber
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
