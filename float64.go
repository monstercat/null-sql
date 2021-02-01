package pgnull

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type NullFloat sql.NullFloat64

func (f NullFloat) MarshalJSON() ([]byte, error) {
	if f.Valid {
		return json.Marshal(f.Float64)
	}
	return json.Marshal(nil)
}

func (f *NullFloat) UnmarshalJSON(bt []byte) error {
	xyz := string(bt)
	if xyz == "null" {
		f.Float64 = 0
		f.Valid = false
		return nil
	}
	v, err := strconv.ParseFloat(xyz, 64)
	if err != nil {
		f.Float64 = 0
		f.Valid = false
		return err
	}
	f.Float64 = v
	f.Valid = true
	return nil
}

func (f *NullFloat) Scan(value interface{}) error {
	switch v := value.(type) {
	case float64:
		f.Float64 = v
		f.Valid = true
	case float32:
		f.Float64 = float64(v)
		f.Valid = true
	case int64:
		f.Float64 = float64(v)
		f.Valid = true
	case int32:
		f.Float64 = float64(v)
		f.Valid = true
	case int:
		f.Float64 = float64(v)
		f.Valid = true
	}
	return nil
}

func (f NullFloat) Value() (driver.Value, error) {
	if !f.Valid {
		return nil, nil
	}
	return f.Float64, nil
}

func NewNullFloat(a float64) NullFloat {
	return NullFloat{a, true}
}

func NullFloatIsEqual(a, b NullFloat) bool {
	if !a.Valid || !b.Valid {
		return a.Valid == b.Valid
	}
	return a.Float64 == b.Float64
}