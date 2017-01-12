package templates

const (
	QuotesTemplate = `{{range .}}Symbol: {{.Symbol}}
  Last:      {{derefFloat .Last | printf "%.2f" }}
  PrevClose: {{derefFloat .Prevclose | printf "%.2f"}}
  Change:    {{derefFloat .Change | printf "%.2f"}} ({{derefFloat .ChangePercentage | printf "%.2f"}}%)
  Low:       {{derefFloat .Low | printf "%.2f"}}
  High:      {{derefFloat .High | printf "%.2f"}}
  Updated:   {{parseDate .TradeDate}}
	{{- end}}`

	OrderTemplate = `{{if .Symbol -}}
  Preview order:
  {{- else -}}
  Order:
  {{- end}}
  {{- if .ID}}
    ID: {{.ID}}
  {{- end}}
  {{- if .Commission}}
    Commission: {{.Commission}}
  {{- end}}
  {{- if .Cost}}
    Cost: {{.Cost}}
  {{- end}}
  {{- if .ExtendedHours}}
    Extended Hours: {{.ExtendedHours}}
  {{- end}}
  {{- if .Fees}}
    Fees: {{.Fees}}
  {{- end}}
  {{- if .MarginChange}}
    Margin Change: {{.MarginChange}}
  {{- end}}
  {{- if .Cost}}
    Cost: {{.Cost}}
  {{- end}}
  {{- if .Status}}
    Status: {{.Status}}
  {{- end}}`

	OrdersTemplate = `Orders:
  {{- range . -}}
  {{- if .ID}}
  ID: {{.ID}}
  {{- end}}
  {{- if .Status}}
    Status: {{.Status}}
  {{- end}}
  {{- end}}`

	PositionsTemplate = `Positions:
  {{- range . -}}
  {{- if .Symbol}}
  Symbol: {{.Symbol}}
  {{- end}}
  {{- if .Status}}
    Status: {{.Status}}
  {{- end}}
  {{- end}}`
)
