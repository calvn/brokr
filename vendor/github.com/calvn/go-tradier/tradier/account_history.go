package tradier

import "fmt"

func (s *AccountService) History(accountId string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/history", accountId)
	return s.AccountRequest(u)
}
