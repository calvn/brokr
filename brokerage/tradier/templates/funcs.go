package templates

import (
	"text/template"
	"time"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"derefFloat": func(data *float64) float64 { return *data },
		"parseDate":  func(rawTime *int64) time.Time { return time.Unix(0, *rawTime*int64(time.Millisecond)) },
	}
}
