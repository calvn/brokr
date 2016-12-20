package tradier

import "fmt"

func (s *WatchlistsService) Get(watchlistId string) (*Watchlist, *Response, error) {
	u := fmt.Sprintf("watchlists/%s", watchlistId)

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
