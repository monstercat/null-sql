package pgnull

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

type NullInt sql.NullInt64

func (i NullInt) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int64)
	}
	return json.Marshal(nil)
}

func (i *NullInt) UnmarshalJSON(bt []byte) error {
	xyz := string(bt)
	if xyz == "null" {
		i.Int64 = 0
		i.Valid = false
		return nil
	}
	v, err := strconv.Atoi(xyz)
	if err != nil {
		i.Int64 = 0
		i.Valid = false
		return err
	}
	i.Int64 = int64(v)
	i.Valid = true
	return nil
}

func (i *NullInt) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		i.Int64 = v
		i.Valid = true
	}
	return nil
}

func (i NullInt) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.Int64, nil
}

func NewNullInt(a int) NullInt {
	return NullInt{int64(a), true}
}

func NullIntIsEqual(a, b NullInt) bool {
	if !a.Valid && !b.Valid {
		return true
	}
	if a.Valid != b.Valid {
		return false
	}
	return a.Int64 == b.Int64
}