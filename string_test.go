package pgnull

import "testing"

func TestNullString_Set(t *testing.T) {
	a := NullString{}
	a.Set("")

	if a.Valid {
		t.Error("Expecting invalid with empty string.")
	}

	a.Set("hihi")
	if !a.Valid {
		t.Error("Expecting valid with non-empty string")
	}

	a.Set("   ")
	if a.Valid {
		t.Error("Expecting invalid with string of just spaces.")
	}
}