package tradier

func (s *OrderService) Preview(accountId string, params *OrderParams) (*Orders, *Response, error) {
	params.Preview = true
	o, resp, err := s.Create(accountId, params)

	return o, resp, err
}
