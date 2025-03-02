package workspaces

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Int32 struct for nullable int32
type Int32 struct {
	sql.NullInt32
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (m *Int32) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Int32)
}

// MarshalJSON ensures correct JSON representation
func (m Int32) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Int32)
}

// UnmarshalYAML handles YAML deserialization
func (m *Int32) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m.Present = true

	var val *int32
	if err := unmarshal(&val); err != nil {
		return err
	}

	if val == nil {
		m.Valid = false
	} else {
		m.Valid = true
		m.Int32 = *val
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (m Int32) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.Int32, nil
}

// GormDataType returns the GORM data type
func (Int32) GormDataType() string {
	return "integer"
}

// NewInt32 creates a new Int32 instance with a value
func NewInt32(value int32) Int32 {
	return Int32{
		NullInt32: sql.NullInt32{Int32: value, Valid: true},
		Present:   true,
	}
}

// NewInt32AutoNull creates an Int32 instance, setting null when the input is explicitly "null"
func NewInt32AutoNull(value string) Int32 {
	if value == "null" {
		return NewInt32Null()
	}

	var intValue int32
	if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
		return NewInt32(intValue)
	}

	return NewInt32Null()
}

// NewInt32Null creates a null Int32 instance
func NewInt32Null() Int32 {
	return Int32{
		NullInt32: sql.NullInt32{Valid: false},
		Present:   true,
	}
}
