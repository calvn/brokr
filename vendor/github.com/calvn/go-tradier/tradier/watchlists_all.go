package tradier

func (s *WatchlistsService) All() (*Watchlists, *Response, error) {
	req, err := s.client.NewRequest("GET", "watchlists", nil)
	if err != nil {
		return nil, nil, err
	}

	wl := &Watchlists{}

	resp, err := s.client.Do(req, wl)
	if err != nil {
		return nil, resp, err
	}

	return wl, resp, nil
}
