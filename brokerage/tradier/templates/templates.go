package templates

const (
	QuotesTemplate = `{{$list := .}}{{range $i, $e := . -}}
Symbol: {{.Symbol}}
  Last:      {{derefFloat $e.Last | printf "%.2f" }}
  PrevClose: {{derefFloat $e.Prevclose | printf "%.2f"}}
  Change:    {{derefFloat $e.Change | printf "%.2f"}} ({{derefFloat $e.ChangePercentage | printf "%.2f"}}%)
  Low:       {{derefFloat $e.Low | printf "%.2f"}}
  High:      {{derefFloat $e.High | printf "%.2f"}}
  Updated:   {{parseDate $e.TradeDate}}
{{- if and (ne (add $i 1) (len $list)) (ne (len $list) 1)}}
{{- println}}
{{- end}}
{{- end}}`

	OrderTemplate = `{{rpad "ID" 10}}{{rpad "SYMBOL" 10}}{{rpad "SIDE" 10}}{{rpad "TYPE" 10}}{{rpad "DURATION" 10}}{{rpad "STATUS" 10}}FILLED/QTY
{{rpad (stringify .ID) 10}}{{rpad (stringify .Symbol) 10}}{{rpad (stringify .Side) 10}}{{rpad (stringify .Type) 10}}{{rpad (stringify .Duration) 10}}{{rpad (stringify .Status) 10}}{{.Filled}}/{{.Quantity}}`

	OrderPreviewTemplate = `{{rpad "SYMBOL" 10}}{{rpad "SIDE" 10}}{{rpad "TYPE" 10}}{{rpad "DURATION" 10}}{{rpad "COMM./COST" 15}}EST. TOTAL
{{rpad (stringify .Symbol) 10}}{{rpad (stringify .Side) 10}}{{rpad (stringify .Type) 10}}{{rpad (stringify .Duration) 10}}{{rpad (printf "%s/%s" (stringify .Commission) (stringify .OrderCost)) 15}}{{stringify .Cost}}`

	OrdersTemplate = `{{rpad "ID" 10}}{{rpad "SYMBOL" 10}}{{rpad "SIDE" 10}}{{rpad "TYPE" 10}}{{rpad "DURATION" 10}}{{rpad "STATUS" 10}}FILLED/QTY
{{- range .}}
{{rpad (stringify .ID) 10}}{{rpad (stringify .Symbol) 10}}{{rpad (stringify .Side) 10}}{{rpad (stringify .Type) 10}}{{rpad (stringify .Duration) 10}}{{rpad (stringify .Status) 10}}{{.Filled}}/{{.Quantity}}
{{- end}}`

	PositionsTemplate = `{{rpad "ID" 10}}{{rpad "SYMBOL" 10}}{{rpad "QUANTITY" 10}}{{rpad "LAST" 10}}{{rpad "CHANGE (%)" 18}}{{rpad "VALUE" 12}}{{rpad "BASIS" 12}}GL (%)
{{- range .}}
{{rpad (stringify .ID) 10}}{{rpad (stringify .Symbol) 10}}{{rpad (stringify .Quantity) 10}}{{rpad (stringify .Last) 10}}{{rpad (printf "%s (%s%%)" (stringify .Change) (stringify .ChangePercentage)) 18}}{{rpad (stringify .Value) 12}}{{rpad (stringify .CostBasis) 12}}{{.GainLoss}} ({{( stringify .GainLossPercentage)}}%)
{{- end}}`
)
