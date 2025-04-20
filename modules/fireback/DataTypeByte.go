package fireback

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// Byte struct for nullable byte (i.e., []byte)
type Byte struct {
	sql.NullByte
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (m *Byte) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Byte)
}

// MarshalJSON ensures correct JSON representation
func (m Byte) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Byte)
}

// UnmarshalYAML handles YAML deserialization
func (m *Byte) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m.Present = true

	var val *byte
	if err := unmarshal(&val); err != nil {
		return err
	}

	if val == nil {
		m.Valid = false
	} else {
		m.Valid = true
		m.Byte = *val
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (m Byte) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.Byte, nil
}

// GormDataType returns the GORM data type
func (Byte) GormDataType() string {
	return "byte"
}

// NewByte creates a new Byte instance with a value
func NewByte(value byte) Byte {
	return Byte{
		NullByte: sql.NullByte{Byte: value, Valid: true},
		Present:  true,
	}
}

// NewByteAutoNull creates a Byte instance, setting null when the input is explicitly "null"
func NewByteAutoNull(value string) Byte {
	if value == "null" {
		return NewByteNull()
	}

	var byteValue byte
	if _, err := fmt.Sscanf(value, "%c", &byteValue); err == nil {
		return NewByte(byteValue)
	}

	return NewByteNull()
}

// NewByteNull creates a null Byte instance
func NewByteNull() Byte {
	return Byte{
		NullByte: sql.NullByte{Valid: false},
		Present:  true,
	}
}
