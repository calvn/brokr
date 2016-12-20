package tradier

import "fmt"

func (s *AccountService) Orders(accountId string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountId)
	return s.AccountRequest(u)
}
