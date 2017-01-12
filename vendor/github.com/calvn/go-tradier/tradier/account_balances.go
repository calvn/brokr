package tradier

import "fmt"

// Balances returns the account's balance for clientID.
func (s *AccountService) Balances(accountID string) (*Balances, *Response, error) {
	u := fmt.Sprintf("accounts/%s/balances", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	b := &Balances{}

	resp, err := s.client.Do(req, b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, nil
}
