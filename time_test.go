package pgnull

import (
	"encoding/json"
	"testing"
)

func TestNullTime_UnmarshalJSON(t *testing.T) {
	var A struct {
		A NullTime
	}

	nullStr := `{"A": null}`
	if err := json.Unmarshal([]byte(nullStr), &A); err != nil {
		t.Fatal(err)
	}
	if A.A.Valid {
		t.Error("should be invalid time")
	}

	withZeroDate := `{"A":"0001-01-01T00:00:00Z"}`
	if err := json.Unmarshal([]byte(withZeroDate), &A); err != nil {
		t.Fatal(err)
	}
	if A.A.Valid {
		t.Error("should be invalid time")
	}

	withDate := `{"A":"2001-01-01T00:00:00Z"}`
	if err := json.Unmarshal([]byte(withDate), &A); err != nil {
		t.Fatal(err)
	}
	if !A.A.Valid {
		t.Error("should be valid time")
	}
}