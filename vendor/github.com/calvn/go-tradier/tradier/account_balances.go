package tradier

import "fmt"

// Balances returns the account's balance for clientID.
func (s *AccountService) Balances(accountID string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/balances", accountID)
	return s.accountRequest(u)
}
