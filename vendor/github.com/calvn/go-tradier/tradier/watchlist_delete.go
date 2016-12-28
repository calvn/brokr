package tradier

import "fmt"

// Delete sends an watchlist deletion request.
func (s *WatchlistsService) Delete(watchlistID string) (*Watchlists, *Response, error) {
	u := fmt.Sprintf("watclists/%s", watchlistID)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, nil, err
	}

	w := &Watchlists{}

	resp, err := s.client.Do(req, w)
	if err != nil {
		return nil, resp, err
	}

	return w, resp, nil
}
