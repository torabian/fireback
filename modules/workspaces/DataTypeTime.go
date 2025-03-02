package workspaces

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Time struct for nullable time
type Time struct {
	sql.NullTime
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (m *Time) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Time)
}

// MarshalJSON ensures correct JSON representation
func (m Time) MarshalJSON() ([]byte, error) {
	if !m.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(m.Time.String())
}

// UnmarshalYAML handles YAML deserialization
func (m *Time) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m.Present = true

	var val *time.Time
	if err := unmarshal(&val); err != nil {
		return err
	}

	if val == nil {
		m.Valid = false
	} else {
		m.Valid = true
		m.Time = *val
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (m Time) MarshalYAML() (interface{}, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.Time.String(), nil
}

// GormDataType returns the GORM data type
func (Time) GormDataType() string {
	return "datetime"
}

// NewTime creates a new Time instance with a value
func NewTime(value time.Time) Time {
	return Time{
		NullTime: sql.NullTime{Time: value, Valid: true},
		Present:  true,
	}
}

// NewTimeNull creates a null Time instance
func NewTimeNull() Time {
	return Time{
		NullTime: sql.NullTime{Valid: false},
		Present:  true,
	}
}
