package tradier

import "encoding/json"

type WatchlistsService service

type Watchlists []*Watchlist

type Watchlist struct {
	Name      *string `json:"name,omitempty"`
	ID        *string `json:"id,omitempty"`
	PublicID  *string `json:"public_id,omitempty"`
	Items     *Items  `json:"items,omitempty"`
	unwrapped bool    // Used internally
}

type watchlist Watchlist

type Items struct {
	Item []*WatchlistItem `json:"item,omitempty"`
}

type items Items

type WatchlistItem struct {
	Symbol *string `json:"symbol,omitempty"`
	ID     *string `json:"id,omitempty"`
}

type watchlistItem struct {
	*WatchlistItem `json:"item,omitempty"`
}

func (i *Items) UnmarshalJSON(b []byte) error {
	itemsStr := ""
	itemsObj := items{}
	itemObj := watchlistItem{}

	// If items is a string, i.e. "null"
	if err := json.Unmarshal(b, &itemsStr); err == nil {
		return nil
	}

	// If itemsr is a JSON array
	if err := json.Unmarshal(b, &itemsObj); err == nil {
		*i = Items(itemsObj)
		return nil
	}

	// If items is an object
	if err := json.Unmarshal(b, &itemObj); err == nil {
		obj := WatchlistItem(*itemObj.WatchlistItem)
		var slice []*WatchlistItem
		slice = append(slice, &obj)
		*i = Items{Item: slice}
		return nil
	}

	return nil
}

func (i *Items) MarshalJSON() ([]byte, error) {
	if len(i.Item) == 0 {
		return json.Marshal("null")
	}

	if len(i.Item) == 1 {
		return json.Marshal(map[string]interface{}{
			"items": i.Item[0],
		})
	}

	return json.Marshal(*i)
}

// Unmarshal json into Watchlist object
func (w *Watchlist) UnmarshalJSON(b []byte) error {
	var wlc struct {
		*watchlist `json:"watchlist,omitempty"`
	}
	wlObj := watchlist{}

	// If wrapped in watchlist object
	if err := json.Unmarshal(b, &wlc); err == nil {
		if wlc.watchlist != nil {
			*w = Watchlist(*wlc.watchlist)
			return nil
		}
	}

	// If not wrapped in anything
	if err := json.Unmarshal(b, &wlObj); err == nil {
		*w = Watchlist(wlObj)
		return nil
	}

	return nil
}

func (w *Watchlist) MarshalJSON() ([]byte, error) {
	if w.unwrapped {
		return json.Marshal(*w)
	}

	return json.Marshal(map[string]interface{}{
		"watchlist": *w,
	})
}

func (w *Watchlists) UnmarshalJSON(b []byte) error {
	var wlc struct {
		W struct {
			W []*Watchlist `json:"watchlist,omitempty"`
		} `json:"watchlists,omitempty"`
	}
	var wlObj struct {
		W struct {
			W *Watchlist `json:"watchlist,omitempty"`
		} `json:"watchlists,omitempty"`
	}
	var wlNull string

	// If watchlist is null
	if err := json.Unmarshal(b, &wlNull); err == nil {
		return nil
	}

	// If watchlist is a JSON array
	if err := json.Unmarshal(b, &wlc); err == nil {
		*w = wlc.W.W
		return nil
	}
	// If watchlist is a single object
	if err := json.Unmarshal(b, &wlObj); err == nil {
		wl := make([]*Watchlist, 0)
		wl = append(wl, wlObj.W.W)
		*w = Watchlists(wl)
		return nil
	}

	return nil
}

func (w *Watchlists) MarshalJSON() ([]byte, error) {
	// Set wrapped to true to marshal differently
	for _, wl := range *w {
		wl.unwrapped = true
	}

	// If []Watchlist is empty
	if len(*w) == 0 {
		return json.Marshal(map[string]interface{}{
			"watchlists": "null",
		})
	}

	// If []Watchlist is size 1, return first and only object
	if len(*w) == 1 {
		wl := *w
		return json.Marshal(map[string]interface{}{
			"watchlists": map[string]interface{}{
				"watchlist": wl[0],
			},
		})
	}

	// Otherwhise marshal normally
	return json.Marshal(map[string]interface{}{
		"watchlists": map[string]interface{}{
			"watchlist": *w,
		},
	})
}
