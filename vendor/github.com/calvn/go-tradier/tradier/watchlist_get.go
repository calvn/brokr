package tradier

import "fmt"

// Get returns the watchlist for watchlistID.
func (s *WatchlistsService) Get(watchlistID string) (*Watchlist, *Response, error) {
	u := fmt.Sprintf("watchlists/%s", watchlistID)

	req, err := s.client.NewRequest("GET", u, nil)
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
