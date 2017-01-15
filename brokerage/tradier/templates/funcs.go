package templates

import (
	"fmt"
	"reflect"
	"strconv"
	"text/template"
	"time"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"add":        func(a, b int) int { return a + b },
		"subtract":   func(a, b float64) float64 { return a - b },
		"rpad":       rpad,
		"derefFloat": func(data *float64) float64 { return *data },
		"stringify":  stringify,
		"parseDate":  func(rawTime *int64) time.Time { return time.Unix(0, *rawTime*int64(time.Millisecond)) },
	}
}

func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

func stringify(v interface{}) string {
	if reflect.ValueOf(v).Kind() == reflect.Ptr {
		rv := reflect.ValueOf(v).Elem()
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(rv.Int(), 10)
		case reflect.String:
			return rv.String()
		}
	}
	return ""
}
