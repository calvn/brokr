package tradier

func (s *UserService) GainLoss() (*User, *Response, error) {
	return s.UserRequest("user/gainloss")
}
