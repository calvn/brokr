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
	var eWrapper struct {
		E struct {
			E []string `json:"error"`
		} `json:"errors"`
	}

	// Allows Errors to be unmarshalled without being wrapped in a parent struct
	if err := json.Unmarshal(b, &eWrapper); err == nil {
		e.AppendStrings(eWrapper.E.E)
		return nil
	}

	return nil
	// return json.Unmarshal(b, e)
}

// MarshalJSON marshals Errors into its JSON representation.
func (e *Errors) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"errors": *e,
	})
}
