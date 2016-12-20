package brokr

func (r *Runner) GetQuotes(symbols []string) error {
	err := (*r.brokerage).GetQuotes(symbols)
	if err != nil {
		return err
	}

	return nil
}
