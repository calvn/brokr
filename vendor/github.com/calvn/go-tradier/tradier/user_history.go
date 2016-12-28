package tradier

// History returns the user's trading history for all accounts.
func (s *UserService) History() (*User, *Response, error) {
	return s.userRequest("user/history")
}
