package fireback

import (
	"database/sql"
	"encoding/json"
)

// Bool struct for nullable boolean
type Bool struct {
	sql.NullBool
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (m *Bool) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Bool)
}

// MarshalJSON ensures correct JSON representation
func (m Bool) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Bool)
}

// UnmarshalYAML handles YAML deserialization
func (m *Bool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m.Present = true

	var val *bool
	if err := unmarshal(&val); err != nil {
		return err
	}

	if val == nil {
		m.Valid = false
	} else {
		m.Valid = true
		m.Bool = *val
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (m Bool) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.Bool, nil
}

// GormDataType returns the GORM data type
func (Bool) GormDataType() string {
	return "boolean"
}

// NewBool creates a new Bool instance with a value
func NewBool(value bool) Bool {
	return Bool{
		NullBool: sql.NullBool{Bool: value, Valid: true},
		Present:  true,
	}
}

// NewBoolAutoNull creates a Bool instance, setting null when the input is explicitly "null"
func NewBoolAutoNull(value string) Bool {
	if value == "null" {
		return NewBoolNull()
	}

	if value == "true" {
		return NewBool(true)
	} else if value == "false" {
		return NewBool(false)
	}

	return NewBoolNull()
}

// NewBoolNull creates a null Bool instance
func NewBoolNull() Bool {
	return Bool{
		NullBool: sql.NullBool{Valid: false},
		Present:  true,
	}
}
