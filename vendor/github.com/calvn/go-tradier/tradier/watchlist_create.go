package tradier

import "github.com/google/go-querystring/query"

type WatchlistParams struct {
	Name    string   `url:"name"`
	Symbols []string `url:"symbols,omitempty,comma"`
}

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
