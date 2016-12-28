package tradier

// Positions returns the user's positions for all accounts.
func (s *UserService) Positions() (*User, *Response, error) {
	return s.userRequest("user/positions")
}
