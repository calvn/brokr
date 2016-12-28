package tradier

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// Create sends an order creation request. This method supports single-sided orders as well as multileg and combo orders.
func (s *OrderService) Create(accountID string, params *OrderParams) (*Orders, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountID)

	// Populate data
	data, err := query.Values(params)
	if err != nil {
		return nil, nil, err
	}

	uri, err := url.Parse(u)
	if err != nil {
		return nil, nil, err
	}

	uri.RawQuery = data.Encode()

	req, err := s.client.NewRequest("POST", u, data.Encode())
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
