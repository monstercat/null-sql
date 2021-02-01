package pgnull

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type NullTime pq.NullTime

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.Time.Format(time.RFC3339))
	}
	return []byte("null"), nil
}

func (nt *NullTime) UnmarshalJSON(bt []byte) error {
	str := string(bt)

	// If the byte array is the characters for "null"
	// then we should ignore it.
	if str == "null" {
		nt.Valid = false
		nt.Time = time.Time{}
		return nil
	}

	str = str[1 : len(str)-1]

	tm, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	if tm.IsZero() {
		nt.Valid = false
	} else {
		nt.Valid = true
		nt.Time = tm
	}
	return nil
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func NullDateIsEqual(a, b NullTime) bool {
	if !a.Valid && !b.Valid {
		return true
	}

	if a.Valid != b.Valid {
		return false
	}

	return dateIsEqual(a.Time, b.Time.In(a.Time.Location()))
}

func NullTimeIsEqual(a, b NullTime) bool {
	if !a.Valid && !b.Valid {
		return true
	}

	if a.Valid != b.Valid {
		return false
	}

	return a.Time.Unix() == b.Time.Unix()
}

func dateIsEqual(a, b time.Time) bool {
	return a.Year() == b.Year() && a.Month() == b.Month() && a.YearDay() == b.YearDay()
}