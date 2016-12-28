package tradier

import "fmt"

// Orders returns the account's orders for accountID.
func (s *AccountService) Orders(accountID string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountID)
	return s.accountRequest(u)
}
