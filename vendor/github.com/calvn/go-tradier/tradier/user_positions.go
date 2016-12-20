package tradier

func (s *UserService) Positions() (*User, *Response, error) {
	return s.UserRequest("user/positions")
}
