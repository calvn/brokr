package tradier

import "fmt"

// OrderStatus returns the account's order status of orderID for accountID.
func (s *AccountService) OrderStatus(accountID, orderID string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)
	return s.accountRequest(u)
}
