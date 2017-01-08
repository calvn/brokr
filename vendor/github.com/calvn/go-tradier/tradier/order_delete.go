package tradier

import "fmt"

// Delete sends an order deletion/cancellation request.
func (s *OrderService) Delete(accountID, orderID string) (*Order, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	o := &Order{}

	resp, err := s.client.Do(req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}
