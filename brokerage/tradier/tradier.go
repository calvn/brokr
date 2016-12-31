package tradier

import (
	"fmt"
	"net/http"

	"github.com/calvn/go-tradier/tradier"
)

type TradierBrokerage struct {
	client *tradier.Client

	// Account is the account ID that will be used
	Account *string
}

func NewTradierBrokerage(httpClient *http.Client) *TradierBrokerage {
	client := tradier.NewClient(httpClient)

	u, _, err := client.User.Profile()
	if err != nil {
		return nil
	}

	// TODO: Handle the case were profile/account slice is empty
	defaultAccount := u.Profile.Account[0].AccountNumber

	return &TradierBrokerage{
		client:  client,
		Account: defaultAccount,
	}
}

func (b *TradierBrokerage) SwitchAccount(accountID string) error {
	p, _, err := b.client.User.Profile()
	if err != nil {
		return err
	}

	for _, account := range *p.Accounts {
		if *account.AccountNumber == accountID {
			*b.Account = accountID
			return nil
		}
	}

	return fmt.Errorf("No account %s found for user.", accountID)
}

func (b *TradierBrokerage) Name() string {
	return "Tradier"
}
