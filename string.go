package pgnull

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type NullString sql.NullString

func (s *NullString) Set(str string) {
	s.String = str
	s.Valid = strings.TrimSpace(str) != ""
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return json.Marshal(nil)
}

func (s *NullString) UnmarshalJSON(bt []byte) error {
	var str string
	if err := json.Unmarshal(bt, &str); err != nil {
		return err
	}
	if len(str) < 1 {
		s.Valid = false
		return nil
	}
	s.String = str
	s.Valid = true
	return nil
}

// Scan implements the Scanner interface.
func (s *NullString) Scan(value interface{}) error {
	s.String, s.Valid = value.(string)
	if !s.Valid {
		bt, valid := value.([]byte)
		if valid {
			s.String = string(bt)
			s.Valid = true
		}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (s NullString) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.String, nil
}

func NewNullString(a string) NullString {
	return NullString{
		String: a,
		Valid:  true,
	}
}

func NullStringIsEqual(a, b NullString) bool {
	if !a.Valid && !b.Valid {
		return true
	}

	if a.Valid != b.Valid {
		return false
	}

	return a.String == b.String
}