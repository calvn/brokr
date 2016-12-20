package tradier

type UserService service

type User struct {
	Profile  *Profile  `json:"profile,omitempty"`
	Accounts *Accounts `json:"accounts,omitempty"`
}

func (s *UserService) UserRequest(uri string) (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, nil, err
	}

	u := &User{}

	resp, err := s.client.Do(req, u)
	if err != nil {
		return nil, resp, err
	}

	return u, resp, nil
}
