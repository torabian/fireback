package fireback

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Int struct for nullable int
type Int struct {
	sql.NullInt32
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (m *Int) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Int32)
}

// MarshalJSON ensures correct JSON representation
func (m Int) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Int32)
}

// UnmarshalYAML handles YAML deserialization
func (m *Int) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m.Present = true

	var val *int
	if err := unmarshal(&val); err != nil {
		return err
	}

	if val == nil {
		m.Valid = false
	} else {
		m.Valid = true
		m.Int32 = int32(*val)
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (m Int) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return int(m.Int32), nil
}

// GormDataType returns the GORM data type
func (Int) GormDataType() string {
	return "integer"
}

// NewInt creates a new Int instance with a value
func NewInt(value int) Int {
	return Int{
		NullInt32: sql.NullInt32{Int32: int32(value), Valid: true},
		Present:   true,
	}
}

// NewIntAutoNull creates an Int instance, setting null when the input is explicitly "null"
func NewIntAutoNull(value string) Int {
	if value == "null" {
		return NewIntNull()
	}

	var intValue int
	if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
		return NewInt(intValue)
	}

	return NewIntNull()
}

// NewIntNull creates a null Int instance
func NewIntNull() Int {
	return Int{
		NullInt32: sql.NullInt32{Valid: false},
		Present:   true,
	}
}
