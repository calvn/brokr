package tradier

import "fmt"

// GainLoss returns the account's cost basis for accountID.
func (s *AccountService) GainLoss(accountID string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/gainloss", accountID)
	return s.accountRequest(u)
}
