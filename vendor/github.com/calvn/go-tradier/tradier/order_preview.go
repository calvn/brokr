package tradier

// Preview sends an order preview request.
func (s *OrderService) Preview(accountID string, params *OrderParams) (*Orders, *Response, error) {
	params.Preview = true
	o, resp, err := s.Create(accountID, params)

	return o, resp, err
}
