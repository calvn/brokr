package tradier

import "fmt"

// Positions returns the account's positions for accountID.
func (s *AccountService) Positions(accountID string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/positions", accountID)
	return s.accountRequest(u)
}
