package tradier

// Balances returns the user's balances for all accounts.
func (s *UserService) Balances() (*User, *Response, error) {
	return s.userRequest("user/balances")
}
