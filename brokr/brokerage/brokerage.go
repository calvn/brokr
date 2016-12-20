package brokerage

type Brokerage interface {
	Name() string
	GetQuotes([]string) error
}
