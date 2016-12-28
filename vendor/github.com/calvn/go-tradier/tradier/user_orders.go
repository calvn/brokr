package tradier

// Orders returns the user's orders for all accounts.
func (s *UserService) Orders() (*User, *Response, error) {
	return s.userRequest("user/orders")
}
