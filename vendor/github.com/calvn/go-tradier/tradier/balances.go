package tradier

type Balances struct {
	AccountNumber *string `json:"account_number,omitempty"`
	AccountType   *string `json:"account_type,omitempty"`
	Cash          *struct {
		CashAvailable  *float64 `json:"cash_available,omitempty"`
		Sweep          *float64 `json:"sweep,omitempty"`
		UnsettledFunds *float64 `json:"unsettled_funds,omitempty"`
	} `json:"cash,omitempty"`
	ClosePL            *float64 `json:"close_pl,omitempty"`
	CurrentRequirement *float64 `json:"current_requirement,omitempty"`
	DividendBalance    *float64 `json:"dividend_balance,omitempty"`
	Equity             *float64 `json:"equity,omitempty"`
	LongLiquidValue    *float64 `json:"long_liquid_value,omitempty"`
	LongLiquidMarket   *float64 `json:"long_market_value,omitempty"`
	Margin             *struct {
		FedCall           *float64 `json:"fed_call,omitempty"`
		MaintenanceCall   *float64 `json:"maintenance_call,omitempty"`
		OptionBuyingPower *float64 `json:"option_buying_power,omitempty"`
		StockBuyingPower  *float64 `json:"stock_buying_power,omitempty"`
		StockShortValue   *float64 `json:"stock_short_value,omitempty"`
		Sweep             *float64 `json:"sweep,omitempty"`
	} `json:"margin,omitempty"`
	MarketValue       *float64 `json:"market_value,omitempty"`
	NetValue          *float64 `json:"net_value,omitempty"`
	OpenPL            *float64 `json:"open_pl,omitempty"`
	OptionLongValue   *float64 `json:"option_long_value,omitempty"`
	OptionRequirement *float64 `json:"option_requirement,omitempty"`
	OptionShortValue  *float64 `json:"option_short_value,omitempty"`
	PDT               *struct {
		DayTradeBuyingPower *float64 `json:"day_trade_buying_power,omitempty"`
		FedCall             *float64 `json:"fed_call,omitempty"`
		MaintenanceCall     *float64 `json:"maintenance_call,omitempty"`
		OptionBuyingPower   *float64 `json:"option_buying_power,omitempty"`
		StockBuyingPower    *float64 `json:"stock_buying_power,omitempty"`
		StockShortValue     *float64 `json:"stock_short_value,omitempty"`
	} `json:"pdt,omitempty"`
	PendingCase        *float64 `json:"pending_cash,omitempty"`
	PendingOrdersCount *int     `json:"pending_orders_count,omitempty"`
	ShortLiquidValue   *float64 `json:"short_liquid_value,omitempty"`
	ShortMarketValue   *float64 `json:"short_market_value,omitempty"`
	StockLongValue     *float64 `json:"stock_long_value,omitempty"`
	UnclearedFunds     *float64 `json:"uncleared_funds,omitempty"`
	TotalCash          *float64 `json:"total_cash,omitempty"`
	TotalEquity        *float64 `json:"total_equity,omitempty"`
}
