package tradier

import "fmt"

func (b *TradierBrokerage) GetPositions() error {
	account, _, err := b.client.Account.Positions(*b.Account)
	if err != nil {
		return err
	}

	fmt.Println(account.Positions)

	return nil
}
