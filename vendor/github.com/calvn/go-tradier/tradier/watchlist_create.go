package tradier

import "github.com/google/go-querystring/query"

// WatchlistParams specifies the query parameters for handling watchlist creation.
// Refer to https://godoc.org/github.com/google/go-querystring/query for building the struct mapping.
type WatchlistParams struct {
	Name    string   `url:"name"`
	Symbols []string `url:"symbols,omitempty,comma"`
}

// Create sends an watchlist creation request.
func (s *WatchlistsService) Create(name string, params *WatchlistParams) (*Watchlist, *Response, error) {
	// Populate data
	data, err := query.Values(params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", "watchlists", data.Encode())
	if err != nil {
		return nil, nil, err
	}

	wl := &Watchlist{}

	resp, err := s.client.Do(req, wl)
	if err != nil {
		return nil, resp, err
	}

	return wl, resp, nil
}
