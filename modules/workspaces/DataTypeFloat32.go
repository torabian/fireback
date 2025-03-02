package workspaces

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Float32 struct for nullable float32
type Float32 struct {
	sql.NullFloat64
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (m *Float32) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Float64)
}

// MarshalJSON ensures correct JSON representation
func (m Float32) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Float64)
}

// UnmarshalYAML handles YAML deserialization
func (m *Float32) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
func (m Float32) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return float32(m.Float64), nil
}

// GormDataType returns the GORM data type
func (Float32) GormDataType() string {
	return "float"
}

// NewFloat32 creates a new Float32 instance with a value
func NewFloat32(value float32) Float32 {
	return Float32{
		NullFloat64: sql.NullFloat64{Float64: float64(value), Valid: true},
		Present:     true,
	}
}

// NewFloat32AutoNull creates a Float32 instance, setting null when the input is explicitly "null"
func NewFloat32AutoNull(value string) Float32 {
	if value == "null" {
		return NewFloat32Null()
	}

	var floatValue float32
	if _, err := fmt.Sscanf(value, "%f", &floatValue); err == nil {
		return NewFloat32(floatValue)
	}

	return NewFloat32Null()
}

// NewFloat32Null creates a null Float32 instance
func NewFloat32Null() Float32 {
	return Float32{
		NullFloat64: sql.NullFloat64{Valid: false},
		Present:     true,
	}
}
