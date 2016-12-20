package tradier

import "fmt"

func (s *AccountService) Positions(accountId string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/positions", accountId)
	return s.AccountRequest(u)
}
