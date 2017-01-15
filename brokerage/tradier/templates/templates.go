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
