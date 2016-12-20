package tradier

import "fmt"

func (s *WatchlistsService) Delete(watchlistId string) (*Watchlists, *Response, error) {
	u := fmt.Sprintf("watclists/%s", watchlistId)

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
