package structs

import "github.com/calvn/go-tradier/tradier"

// PositionWrapper is wrapper on tradier.Positions with additional
// fields used for positions output.
type PositionWrapper struct {
	*tradier.Position
	Last               *float64
	Change             *float64
	ChangePercentage   *float64
	Value              *float64
	GainLoss           *float64
	GainLossPercentage *float64
}

// PositionsWrapper holds a collection of PositionWrapper.
type PositionsWrapper []*PositionWrapper

// NewPositionWrapper returns *OrderTemplate with the filled amount calculated.
func NewPositionWrapper(position *tradier.Position, quote *tradier.Quote) *PositionWrapper {
	p := &PositionWrapper{Position: position}
	p.Last = quote.Last
	p.Change = quote.Change
	p.ChangePercentage = quote.ChangePercentage
	// TODO: Check that float correctly multiplies
	p.Value = tradier.Float64(*p.Quantity * *quote.Last)
	p.GainLoss = tradier.Float64(*p.Value - *p.CostBasis)
	p.GainLossPercentage = tradier.Float64((*p.GainLoss / *p.CostBasis) * 100.00)

	return p
}

func NewPositionsWrapper(positions *tradier.Positions, quotes *tradier.Quotes) *PositionsWrapper {
	p := make(PositionsWrapper, len(*positions))
	for i, position := range *positions {
		p[i] = NewPositionWrapper(position, (*quotes)[i])
	}

	return &p
}
