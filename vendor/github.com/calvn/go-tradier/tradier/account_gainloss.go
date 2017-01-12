package tradier

import "fmt"

// GainLoss returns the account's cost basis for accountID.
func (s *AccountService) GainLoss(accountID string) (*GainLoss, *Response, error) {
	u := fmt.Sprintf("accounts/%s/gainloss", accountID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	g := &GainLoss{}

	resp, err := s.client.Do(req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}
