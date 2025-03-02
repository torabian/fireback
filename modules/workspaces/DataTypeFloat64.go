package workspaces

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Float64 struct for nullable float64
type Float64 struct {
	sql.NullFloat64
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (m *Float64) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Float64)
}

// MarshalJSON ensures correct JSON representation
func (m Float64) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Float64)
}

// UnmarshalYAML handles YAML deserialization
func (m *Float64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m.Present = true

	var val *float64
	if err := unmarshal(&val); err != nil {
		return err
	}

	if val == nil {
		m.Valid = false
	} else {
		m.Valid = true
		m.Float64 = *val
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (m Float64) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.Float64, nil
}

// GormDataType returns the GORM data type
func (Float64) GormDataType() string {
	return "double"
}

// NewFloat64 creates a new Float64 instance with a value
func NewFloat64(value float64) Float64 {
	return Float64{
		NullFloat64: sql.NullFloat64{Float64: value, Valid: true},
		Present:     true,
	}
}

// NewFloat64AutoNull creates a Float64 instance, setting null when the input is explicitly "null"
func NewFloat64AutoNull(value string) Float64 {
	if value == "null" {
		return NewFloat64Null()
	}

	var floatValue float64
	if _, err := fmt.Sscanf(value, "%f", &floatValue); err == nil {
		return NewFloat64(floatValue)
	}

	return NewFloat64Null()
}

// NewFloat64Null creates a null Float64 instance
func NewFloat64Null() Float64 {
	return Float64{
		NullFloat64: sql.NullFloat64{Valid: false},
		Present:     true,
	}
}
