package errors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Errors hold a collection of errors returned by a response.
type Errors struct {
	Err []error `json:"error"`
}

func (e *Errors) Error() string {
	if len(e.Err) == 1 {
		return fmt.Sprintf("1 error occurred:\n  %s", e.Err[0])
	}

	errors := make([]string, len(e.Err))
	for i, err := range e.Err {
		errors[i] = fmt.Sprintf("  %s", err)
	}

	return fmt.Sprintf("%d errors occurred:\n%s", len(e.Err), strings.Join(errors, "\n"))
}

// AppendStrings appends a slice of strings to *Errors.Err.
func (e *Errors) AppendStrings(errorStrings []string) {
	for _, s := range errorStrings {
		e.Err = append(e.Err, fmt.Errorf(s))
	}
}

// UnmarshalJSON unmarshals errors into Errors object.
func (e *Errors) UnmarshalJSON(b []byte) error {
	var eCol struct {
		E struct {
			E []string `json:"error"`
		} `json:"errors"`
	}
	var eObj struct {
		E struct {
			E string `json:"error,omitempty"`
		} `json:"errors,omitempty"`
	}
	var err error

	// If error is wrapped in an array
	if err = json.Unmarshal(b, &eCol); err == nil {
		e.AppendStrings(eCol.E.E)
		return nil
	}

	// If error is wrapped in an object
	if err = json.Unmarshal(b, &eObj); err == nil {
		e.AppendStrings([]string{eObj.E.E})
		return nil
	}

	return err
}

// MarshalJSON marshals Errors into its JSON representation.
func (e *Errors) MarshalJSON() ([]byte, error) {
	if len(e.Err) == 1 {
		return json.Marshal(map[string]interface{}{
			"errors": map[string]interface{}{
				"error": (*e).Err[0],
			},
		})
	}

	return json.Marshal(map[string]interface{}{
		"errors": map[string]interface{}{
			"error": *e,
		},
	})
}
