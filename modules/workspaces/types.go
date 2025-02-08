package workspaces

import (
	"database/sql"
	"encoding/json"
)

// Bool struct for nullable boolean
type Bool struct {
	sql.NullBool
	Present bool // Whether the field was present in JSON
}

func (m *Bool) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Bool)
}

// Int struct for nullable int
type Int struct {
	sql.NullInt32
	Present bool // Whether the field was present in JSON
}

func (m *Int) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Int32)
}

// Int32 struct for nullable int32
type Int32 struct {
	sql.NullInt32
	Present bool // Whether the field was present in JSON
}

func (m *Int32) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Int32)
}

// Int64 struct for nullable int64
type Int64 struct {
	sql.NullInt64
	Present bool // Whether the field was present in JSON
}

func (m *Int64) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Int64)
}

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

// Time struct for nullable time
type Time struct {
	sql.NullTime
	Present bool // Whether the field was present in JSON
}

func (m *Time) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Time)
}

// Nullable Float32
type Float32 struct {
	sql.NullFloat64
	Present bool
}

func (m *Float32) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Float64)
}

// Nullable Float64
type Float64 struct {
	sql.NullFloat64
	Present bool
}

func (m *Float64) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.Float64)
}

// Nullable Byte
type Byte struct {
	sql.NullString
	Present bool
}

func (m *Byte) UnmarshalJSON(data []byte) error {
	m.Present = true
	if string(data) == "null" {
		m.Valid = false
		return nil
	}
	m.Valid = true
	return json.Unmarshal(data, &m.String)
}

func NewInt32(value int32) *Int32 {
	return &Int32{
		NullInt32: sql.NullInt32{Int32: value, Valid: true},
		Present:   true,
	}
}

func NewInt32Null() *Int32 {
	return &Int32{
		NullInt32: sql.NullInt32{Valid: false},
		Present:   true,
	}
}

func NewInt64(value int64) *Int64 {
	return &Int64{
		NullInt64: sql.NullInt64{Int64: value, Valid: true},
		Present:   true,
	}
}

func NewInt64Null() *Int64 {
	return &Int64{
		NullInt64: sql.NullInt64{Valid: false},
		Present:   true,
	}
}

func NewString(value string) String {
	return String{
		NullString: sql.NullString{String: value, Valid: true},
		Present:    true,
	}
}

func NewStringNull() String {
	return String{
		NullString: sql.NullString{Valid: false},
		Present:    true,
	}
}
