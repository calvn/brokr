package tradier

import "fmt"

func (s *AccountService) Balances(accountId string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/balances", accountId)
	return s.AccountRequest(u)
}
