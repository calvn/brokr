package tradier

import "fmt"

// History returns the account's trading history for accountID.
func (s *AccountService) History(accountID string) (*History, *Response, error) {
	u := fmt.Sprintf("accounts/%s/history", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	h := &History{}

	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, nil
}
