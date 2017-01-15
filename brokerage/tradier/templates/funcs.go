package templates

import (
	"text/template"
	"time"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"add":        func(a, b int) int { return a + b },
		"derefFloat": func(data *float64) float64 { return *data },
		"parseDate":  func(rawTime *int64) time.Time { return time.Unix(0, *rawTime*int64(time.Millisecond)) },
	}
}
