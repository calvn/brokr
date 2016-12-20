package tradier

func (s *UserService) Balances() (*User, *Response, error) {
	return s.UserRequest("user/balances")
}
