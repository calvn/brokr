package tradier

import "fmt"

func (b *Brokerage) GetPositions() error {
	account, _, err := b.client.Account.Positions(*b.AccountID)
	if err != nil {
		return err
	}

	fmt.Println(account.Positions)

	return nil
}
