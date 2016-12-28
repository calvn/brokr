package tradier

// Profile represents the profile JSON object.
type Profile struct {
	Account []*Account `json:"account,omitempty"`
	ID      *string    `json:"id,omitempty"`
	Name    *string    `json:"name,omitempty"`
}
