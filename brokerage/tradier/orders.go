package tradier

// TODO: Move templates into /templates

var orderTemplate = `{{if .Symbol -}}
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

var ordersTemplate = `Orders:
{{- range . -}}
{{- if .ID}}
ID: {{.ID}}
{{- end}}
{{- if .Status}}
  Status: {{.Status}}
{{- end}}
{{- end}}`

var positionsTemplate = `Positions:
{{- range . -}}
{{- if .Symbol}}
Symbol: {{.Symbol}}
{{- end}}
{{- if .Status}}
  Status: {{.Status}}
{{- end}}
{{- end}}`
