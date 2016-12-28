package tradier

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

var accountsJSONSingle = []byte(`{
	"accounts": {
		"account": {
			"account_number": "6YA05708"
		}
	}
}`)

var accountsJSONArray = []byte(`{
	"accounts": {
		"account": [{
			"account_number": "6YA05708"
		}, {
			"account_number": "6YA05709"
		}]
	}
}`)

var accountsJSONNull = []byte(`{
	"accounts": "null"
}`)

var accountsSingle = &Accounts{
	{
		AccountNumber: String("6YA05708"),
	},
}

var accountsArray = &Accounts{
	{
		AccountNumber: String("6YA05708"),
	},
	{
		AccountNumber: String("6YA05709"),
	},
}

var accountsNull = &Accounts{}

func TestAccounts_UnmarshalJSON_Single(t *testing.T) {
	want := accountsSingle

	got := &Accounts{}
	err := json.Unmarshal(accountsJSONSingle, got)
	if err != nil {
		t.Errorf("Accounts.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}

func TestAccounts_MarshalJSON_Single(t *testing.T) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, accountsJSONSingle)
	want := buf.Bytes()
	if err != nil {
		t.Error(err)
	}

	got, err := json.Marshal(accountsSingle)
	if err != nil {
		t.Errorf("Accounts.MarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func TestAccounts_UnmarshalJSON_Array(t *testing.T) {
	want := accountsArray

	got := &Accounts{}
	err := json.Unmarshal(accountsJSONArray, got)
	if err != nil {
		t.Errorf("Accounts.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}

func TestAccounts_MarshalJSON_Array(t *testing.T) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, accountsJSONArray)
	want := buf.Bytes()
	if err != nil {
		t.Error(err)
	}

	got, err := json.Marshal(accountsArray)
	if err != nil {
		t.Errorf("Accounts.MarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func TestAccounts_UnmarshalJSON_Null(t *testing.T) {
	want := accountsNull

	got := &Accounts{}
	err := json.Unmarshal(accountsJSONNull, got)
	if err != nil {
		t.Errorf("Accounts.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}

func TestAccounts_MarshalJSON_Null(t *testing.T) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, accountsJSONNull)
	want := buf.Bytes()
	if err != nil {
		t.Error(err)
	}

	got, err := json.Marshal(&accountsNull)
	if err != nil {
		t.Errorf("Accounts.MarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s want: %s", got, want)
	}
}
