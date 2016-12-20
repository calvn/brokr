package tradier

import "fmt"

func (s *AccountService) GainLoss(accountId string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/gainloss", accountId)
	return s.AccountRequest(u)
}
