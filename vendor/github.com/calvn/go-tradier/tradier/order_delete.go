package tradier

import "fmt"

func (s *OrderService) Delete(accountId, orderId string) (*Orders, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountId, orderId)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	o := &Orders{}

	resp, err := s.client.Do(req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}
