package tradier

func (s *UserService) History() (*User, *Response, error) {
	return s.UserRequest("user/history")
}
