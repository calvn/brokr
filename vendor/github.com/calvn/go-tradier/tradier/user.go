package tradier

import "encoding/json"

// UserService handles routes related to user inquiry
// from the Tradier API.
type UserService service

// User represents the user JSON object.
type User struct {
	// Profile is specific to users/profile
	Profile *Profile `json:"profile,omitempty"`

	Accounts *Accounts `json:"accounts,omitempty"`
}

func (s *UserService) userRequest(uri string) (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, nil, err
	}

	u := &User{}

	resp, err := s.client.Do(req, u)
	if err != nil {
		return nil, resp, err
	}

	return u, resp, nil
}

// MarshalJSON marshals Accounts into JSON
func (u *User) MarshalJSON() ([]byte, error) {
	if u.Profile == nil {
		return json.Marshal(u.Accounts)
	}

	return json.Marshal(u)
}
