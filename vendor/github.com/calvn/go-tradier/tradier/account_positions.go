package tradier

import "fmt"

// Positions returns the account's positions for accountID.
func (s *AccountService) Positions(accountID string) (*Positions, *Response, error) {
	u := fmt.Sprintf("accounts/%s/positions", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	p := &Positions{}

	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}
