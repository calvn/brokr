package brokr

type Brokerage interface {
	GetQuotes([]string) error
}
