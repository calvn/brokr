package tradier

import (
	"fmt"
	"net/url"
	"strings"
)

// Quotes returns the quotes for a set of symbols.
func (s *MarketsService) Quotes(symbols []string) (*Quotes, *Response, error) {
	symbls := strings.Join(symbols, ",")

	v := url.Values{}
	v.Add("symbols", symbls)

	u := fmt.Sprintf("markets/quotes?%s", v.Encode())

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	q := &Quotes{}

	resp, err := s.client.Do(req, q)
	if err != nil {
		return nil, resp, err
	}

	return q, resp, nil
}
