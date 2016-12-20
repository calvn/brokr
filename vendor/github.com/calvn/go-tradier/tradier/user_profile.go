package tradier

func (s *UserService) Profile() (*User, *Response, error) {
	return s.UserRequest("user/profile")
}
