package tradier

// GainLoss returns the user's cost basis for all accounts.
func (s *UserService) GainLoss() (*User, *Response, error) {
	return s.userRequest("user/gainloss")
}
