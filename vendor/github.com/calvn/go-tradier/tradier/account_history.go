package tradier

import "fmt"

// History returns the account's trading history for accountID.
func (s *AccountService) History(accountID string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/history", accountID)
	return s.accountRequest(u)
}
