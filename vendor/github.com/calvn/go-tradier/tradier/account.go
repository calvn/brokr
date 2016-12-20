package tradier

type AccountService service

func (s *AccountService) AccountRequest(uri string) (*Account, *Response, error) {
	req, err := s.client.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, nil, err
	}

	a := &Account{}

	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}
