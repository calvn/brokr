package tradier

// Profile returns the user's profile information. This includes basic information for all accounts under that user.
func (s *UserService) Profile() (*User, *Response, error) {
	return s.userRequest("user/profile")
}
