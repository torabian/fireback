package workspaces

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"unicode"

	duration "github.com/torabian/fireback/modules/workspaces/external"
)

// Duration struct for nullable duration
type Duration struct {
	sql.NullInt64
	Present bool // Whether the field was present in JSON or YAML
}

// UnmarshalJSON handles JSON deserialization
func (d *Duration) UnmarshalJSON(data []byte) error {
	d.Present = true
	if string(data) == "null" {
		d.Valid = false
		return nil
	}

	// Parse the ISO 8601 duration string into seconds
	var isoDuration string
	if err := json.Unmarshal(data, &isoDuration); err != nil {
		return err
	}

	parsedDuration := parseStringToSeconds(isoDuration)
	d.Valid = true
	d.Int64 = parsedDuration
	return nil
}

// MarshalJSON ensures correct JSON representation
func (d Duration) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte("null"), nil
	}

	// Convert the stored seconds back to an ISO 8601 duration string
	isoDuration := formatISODuration(d.Int64)
	return json.Marshal(isoDuration)
}

// UnmarshalYAML handles YAML deserialization
func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	d.Present = true

	var val *string
	if err := unmarshal(&val); err != nil {
		return err
	}

	if val == nil {
		d.Valid = false
	} else {
		parsedDuration := parseStringToSeconds(*val)
		d.Valid = true
		d.Int64 = parsedDuration
	}
	return nil
}

// MarshalYAML ensures correct YAML representation
func (d Duration) MarshalYAML() (interface{}, error) {
	if !d.Valid {
		return nil, nil
	}

	// Convert the stored seconds back to an ISO 8601 duration string
	isoDuration := formatISODuration(d.Int64)
	return isoDuration, nil
}

// GormDataType returns the GORM data type
func (Duration) GormDataType() string {
	return "bigint"
}

// NewDuration creates a new Duration instance with a value
func NewDuration(value int64) Duration {
	return Duration{
		NullInt64: sql.NullInt64{Int64: value, Valid: true},
		Present:   true,
	}
}

// NewDurationAutoNull creates a Duration instance, setting null when the input is explicitly "null"
func NewDurationAutoNull(value string) Duration {
	if value == "null" {
		return NewDurationNull()
	}

	parsedDuration := parseStringToSeconds(value)
	return NewDuration(parsedDuration)
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// ToInt64 casts a string to int64 if it's numeric
func ToInt64(s string) (int64, error) {
	if IsNumeric(s) {
		return strconv.ParseInt(s, 10, 64)
	}
	return 0, fmt.Errorf("not a valid number")
}

// NewDurationNull creates a null Duration instance
func NewDurationNull() Duration {
	return Duration{
		NullInt64: sql.NullInt64{Valid: false},
		Present:   true,
	}
}

func parseStringToSeconds(input string) int64 {
	if IsNumeric(input) {
		if v, err := ToInt64(input); err == nil {
			return v * 1_000_000_000
		}
	}

	if dur, err := duration.Parse(input); err == nil {
		return int64(dur.ToTimeDuration())
	} else {
		fmt.Println("Error", err)
	}

	return 0
}

func formatISODuration(seconds int64) string {
	return duration.FromTimeDuration(time.Duration(seconds)).String()
}
