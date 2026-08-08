package complexes

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// DefaultLocale is the locale TString falls back to: whenever it's built or
// parsed from a plain string instead of a {locale: value} object, and
// whenever Get()/String() are asked for a locale that isn't present.
const DefaultLocale = "en"

// TString ("translatable string") stores a piece of text as a
// locale -> value map, e.g. {"en": "Hello", "fa": "سلام"}, and is stored in
// the database as a single JSON column, the same way the JSON type in this
// package is - no separate table or per-language columns needed to hold
// every translation of a string.
//
// Any place a TString is read from JSON, YAML, or plain Go/CLI text as a
// bare string ("Hello") rather than an object, it's treated as the
// DefaultLocale's value ({"en": "Hello"}) - so callers that don't care
// about translations can keep passing plain strings.
type TString map[string]string

// NewTString builds a TString with just DefaultLocale set - the common case
// of "I don't have translations yet, just a string".
func NewTString(value string) TString {
	return TString{DefaultLocale: value}
}

// TStringFrom wraps an existing locale->value map as a TString.
func TStringFrom(values map[string]string) TString {
	return TString(values)
}

// Get returns the value for locale, falling back to DefaultLocale, then to
// any single value present, then "".
func (t TString) Get(locale string) string {
	if t == nil {
		return ""
	}
	if v, ok := t[locale]; ok {
		return v
	}
	if v, ok := t[DefaultLocale]; ok {
		return v
	}
	for _, v := range t {
		return v
	}
	return ""
}

// Has reports whether locale has an explicit value set.
func (t TString) Has(locale string) bool {
	if t == nil {
		return false
	}
	_, ok := t[locale]
	return ok
}

// Set sets locale's value, allocating the map if t is nil, and returns t
// for chaining (e.g. title := complexes.NewTString("Hi").Set("fa", "سلام")).
func (t TString) Set(locale string, value string) TString {
	if t == nil {
		t = TString{}
	}
	t[locale] = value
	return t
}

// Locales returns every locale code that has a value set.
func (t TString) Locales() []string {
	locales := make([]string, 0, len(t))
	for locale := range t {
		locales = append(locales, locale)
	}
	return locales
}

// String implements fmt.Stringer, returning Get(DefaultLocale).
func (t TString) String() string {
	return t.Get(DefaultLocale)
}

// parseTString accepts either a bare JSON string (wrapped into
// DefaultLocale) or a {locale: value} object - the same duality
// UnmarshalJSON/UnmarshalYAML/Scan all need, so it's shared here.
func parseTString(data []byte) (TString, error) {
	trimmed := strings.TrimSpace(string(data))
	if len(trimmed) == 0 || trimmed == "null" {
		return nil, nil
	}

	var asString string
	if err := json.Unmarshal(data, &asString); err == nil {
		return TString{DefaultLocale: asString}, nil
	}

	var asMap map[string]string
	if err := json.Unmarshal(data, &asMap); err != nil {
		return nil, fmt.Errorf("TString: expected a string or a {locale: value} object: %w", err)
	}
	return TString(asMap), nil
}

func (t *TString) UnmarshalJSON(data []byte) error {
	parsed, err := parseTString(data)
	if err != nil {
		return err
	}
	*t = parsed
	return nil
}

func (t *TString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var asString string
	if err := unmarshal(&asString); err == nil {
		*t = TString{DefaultLocale: asString}
		return nil
	}

	var asMap map[string]string
	if err := unmarshal(&asMap); err != nil {
		return err
	}
	*t = TString(asMap)
	return nil
}

// UnmarshalText covers plain-text sources (CLI flags, form fields, etc.)
// that only ever hand TString a bare string - always DefaultLocale.
func (t *TString) UnmarshalText(text []byte) error {
	*t = TString{DefaultLocale: string(text)}
	return nil
}

// Value implements driver.Valuer, so gorm stores TString as a JSON object.
func (t TString) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}
	b, err := json.Marshal(map[string]string(t))
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// Scan implements sql.Scanner.
func (t *TString) Scan(value interface{}) error {
	if value == nil {
		*t = nil
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("TString: unsupported db value: ", value))
	}

	parsed, err := parseTString(bytes)
	if err != nil {
		return err
	}
	*t = parsed
	return nil
}

// GormDataType tells gorm this is a JSON-shaped field.
func (TString) GormDataType() string {
	return "json"
}

// GormDBDataType picks the right JSON column type per dialect - same
// mapping as the JSON type in this package.
func (TString) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql", "mariadb":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

// GormValue mirrors JSON.GormValue: MySQL (non-MariaDB) needs the literal
// cast to satisfy its strict JSON column typing; every other dialect is
// happy with the plain string value from Value().
func (t TString) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	data, err := t.Value()
	if err != nil || data == nil {
		return gorm.Expr("NULL")
	}

	str, _ := data.(string)

	switch db.Dialector.Name() {
	case "mysql":
		if v, ok := db.Dialector.(*mysql.Dialector); ok && !strings.Contains(v.ServerVersion, "MariaDB") {
			return gorm.Expr("CAST(? AS JSON)", str)
		}
	}

	return gorm.Expr("?", str)
}
