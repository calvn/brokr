package tradier

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// NOTE: Since this is using OrderParams, we should to some sort of checking or improve on error handling

// Update sends an order update request.
func (s *OrderService) Update(accountID, orderID string, params *OrderParams) (*Order, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountID, orderID)

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

	req, err := s.client.NewRequest("PUT", u, data.Encode())
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
