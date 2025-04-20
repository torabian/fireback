package fireback

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Int64 struct for nullable int64
type Int64 struct {
	sql.NullInt64
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (m *Int64) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Int64)
}

// MarshalJSON ensures correct JSON representation
func (m Int64) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Int64)
}

// UnmarshalYAML handles YAML deserialization
func (m *Int64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m.Present = true

	var val *int64
	if err := unmarshal(&val); err != nil {
		return err
	}

	if val == nil {
		m.Valid = false
	} else {
		m.Valid = true
		m.Int64 = *val
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (m Int64) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.Int64, nil
}

// GormDataType returns the GORM data type
func (Int64) GormDataType() string {
	return "bigint"
}

// NewInt64 creates a new Int64 instance with a value
func NewInt64(value int64) Int64 {
	return Int64{
		NullInt64: sql.NullInt64{Int64: value, Valid: true},
		Present:   true,
	}
}

// NewInt64AutoNull creates an Int64 instance, setting null when the input is explicitly "null"
func NewInt64AutoNull(value string) Int64 {
	if value == "null" {
		return NewInt64Null()
	}

	var intValue int64
	if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
		return NewInt64(intValue)
	}

	return NewInt64Null()
}

// NewInt64Null creates a null Int64 instance
func NewInt64Null() Int64 {
	return Int64{
		NullInt64: sql.NullInt64{Valid: false},
		Present:   true,
	}
}
