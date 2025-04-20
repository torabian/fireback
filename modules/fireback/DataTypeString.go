package fireback

import (
	"database/sql"
	"encoding/json"
)

// String struct for nullable string
type String struct {
	sql.NullString
	Present bool // Whether the field was present in JSON
}

func (m *String) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.String)
}

// MarshalJSON ensures correct JSON representation
func (m String) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.String)
}

// UnmarshalYAML handles YAML deserialization
func (m *String) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m.Present = true

	var str *string
	if err := unmarshal(&str); err != nil {
		return err
	}

	if str == nil {
		m.Valid = false
	} else {
		m.Valid = true
		m.String = *str
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (m String) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.String, nil
}

func (String) GormDataType() string {
	return "string"
}

func NewString(value string) String {
	return String{
		NullString: sql.NullString{String: value, Valid: true},
		Present:    true,
	}
}

func NewStringAutoNull(value string) String {
	if value == "null" {
		return NewStringNull()
	}

	return NewString(value)
}

func NewStringNull() String {
	return String{
		NullString: sql.NullString{Valid: false},
		Present:    true,
	}
}
