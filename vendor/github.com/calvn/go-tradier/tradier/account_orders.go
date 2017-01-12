package tradier

import "fmt"

// Orders returns the account's orders for accountID.
func (s *AccountService) Orders(accountID string) (*Orders, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	o := &Orders{}

	resp, err := s.client.Do(req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}
