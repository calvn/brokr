package tradier

func (s *UserService) Orders() (*User, *Response, error) {
	return s.UserRequest("user/orders")
}
