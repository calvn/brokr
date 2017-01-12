package tradier

import "fmt"

// OrderStatus returns the account's order status of orderID for accountID.
func (s *AccountService) OrderStatus(accountID, orderID string) (*Order, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	o := &Order{}

	resp, err := s.client.Do(req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}
