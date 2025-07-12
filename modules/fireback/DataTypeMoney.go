package fireback

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Money represents a monetary value with currency and nullable support
type Money struct {
	Amount   sql.NullInt64
	Currency sql.NullString
	Present  bool
}

func currencyMultiplier(currency string) int64 {
	return pow10(CurrencyPrecision(currency))
}

func normalizeRawAmount(val float64, currency string) int64 {
	return int64(val * float64(currencyMultiplier(currency)))
}
func (m *Money) parseMoneyInput(input interface{}) error {
	m.Present = true

	switch v := input.(type) {
	case string:
		amount, currency, err := ParseMoneyString(v)
		if err != nil {
			return err
		}
		m.Amount = sql.NullInt64{Int64: amount, Valid: true}
		m.Currency = sql.NullString{String: NormalizeCurrency(currency), Valid: true}
		return nil

	case map[string]interface{}:
		var (
			amountSet   bool
			currencyStr string
			multiplier  int64 = 1
		)

		if rawCurrency, ok := v["currency"].(string); ok && rawCurrency != "" {
			currencyStr = NormalizeCurrency(rawCurrency)
			if currencyStr == "" {
				return fmt.Errorf("currency not recognized")
			}
			multiplier = currencyMultiplier(currencyStr)
			m.Currency = sql.NullString{String: currencyStr, Valid: true}
		} else {
			return fmt.Errorf("currency field is required and must be a non-empty string")
		}

		if rawAmount, ok := v["amount"]; ok {
			switch val := rawAmount.(type) {
			case float64:
				m.Amount = sql.NullInt64{Int64: int64(val * float64(multiplier)), Valid: true}
				amountSet = true
			case int:
				m.Amount = sql.NullInt64{Int64: int64(val) * multiplier, Valid: true}
				amountSet = true
			case int64:
				m.Amount = sql.NullInt64{Int64: val * multiplier, Valid: true}
				amountSet = true
			}
		}

		if !amountSet && !m.Currency.Valid {
			return fmt.Errorf("invalid money object")
		}
		return nil

	default:
		return fmt.Errorf("unsupported input type for money")
	}
}

// JSON unmarshal calls the above
func (m *Money) UnmarshalJSON(data []byte) error {
	// Try object first
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	// If null
	if obj == nil {
		m.Amount.Valid = false
		m.Currency.Valid = false
		return nil
	}

	// Delegate
	return m.parseMoneyInput(obj)
}

// YAML unmarshal calls the same helper
func (m *Money) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var obj interface{}
	if err := unmarshal(&obj); err != nil {
		return err
	}

	// Delegate
	return m.parseMoneyInput(obj)
}

func CurrencyPrecision(currency string) int {
	switch strings.ToUpper(currency) {
	case "BTC":
		return 8
	case "ETH":
		return 10
	case "USD", "EUR", "PLN", "TRY":
		return 2
	case "IRR":
		return 0
	default:
		return 2 // safe fallback
	}
}

func pow10(n int) int64 {
	result := int64(1)
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

type moneyOutput struct {
	Amount    *float64 `json:"amount,omitempty" yaml:"amount,omitempty"`
	Currency  string   `json:"currency,omitempty" yaml:"currency,omitempty"`
	Formatted string   `json:"formatted,omitempty" yaml:"formatted,omitempty"`
}

func (m Money) export(locale string) *moneyOutput {
	if !m.Amount.Valid && !m.Currency.Valid {
		return nil
	}

	precision := CurrencyPrecision(m.Currency.String)
	var amountFloat *float64
	if m.Amount.Valid {
		val := float64(m.Amount.Int64) / float64(pow10(precision))
		amountFloat = &val
	}

	return &moneyOutput{
		Amount:    amountFloat,
		Currency:  m.Currency.String,
		Formatted: m.Format(locale),
	}
}

func (m Money) Formatted() string {
	return m.Format("en-us")
}

func (m Money) exportJSONOrYAML() *moneyOutput {
	if !m.Amount.Valid && !m.Currency.Valid {
		return nil
	}
	return m.export("en-US")
}

func (m Money) MarshalJSON() ([]byte, error) {
	out := m.exportJSONOrYAML()
	if out == nil {
		return []byte("null"), nil
	}
	return json.Marshal(out)
}

func (m Money) MarshalYAML() (interface{}, error) {
	return m.exportJSONOrYAML(), nil
}

// GORM integration - flatten to 2 fields
func (Money) GormDataType() string {
	return "embedded"
}

func (m Money) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	// Not used directly, GORM handles embedded structs
	return clause.Expr{}
}

// Scan implements the sql.Scanner interface
func (m *Money) Scan(value interface{}) error {
	return errors.New("direct scan not supported, use embedded columns")
}

// Value implements the driver.Valuer interface
func (m Money) Value() (driver.Value, error) {
	return nil, errors.New("direct value not supported, use embedded columns")
}

// Helpers
func ifValidInt64Ptr(v sql.NullInt64) *int64 {
	if v.Valid {
		return &v.Int64
	}
	return nil
}
func ifValidStringPtr(v sql.NullString) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}
func ifValidInt64(v sql.NullInt64) interface{} {
	if v.Valid {
		return v.Int64
	}
	return nil
}
func ifValidString(v sql.NullString) interface{} {
	if v.Valid {
		return v.String
	}
	return nil
}

// NewMoney creates a Money instance
func NewMoney(amount int64, currency string) Money {
	return Money{
		Amount:   sql.NullInt64{Int64: amount, Valid: true},
		Currency: sql.NullString{String: currency, Valid: true},
		Present:  true,
	}
}

func NewMoneyAutoNull(value string) Money {
	if value == "null" {
		return NewMoneyNull()
	}

	amount, currency, err := ParseMoneyString(value)

	if err == nil {
		return NewMoney(amount, currency)
	}

	return NewMoney(0, "USD")
}

// NewMoneyNull creates a null Money instance
func NewMoneyNull() Money {
	return Money{
		Amount:   sql.NullInt64{Valid: false},
		Currency: sql.NullString{Valid: false},
		Present:  true,
	}
}

func ParseMoneyString(input string) (int64, string, error) {
	input = strings.ToLower(strings.TrimSpace(input))

	// Remove thousand separators (commas)
	input = strings.ReplaceAll(input, ",", "")

	// Extended regex to match currency symbol/code before or after amount with optional spaces
	// Supports decimals with dot or comma
	regex := regexp.MustCompile(`(?i)^(?:` +
		`(?P<currency1>btc|eth|usd|eur|gbp|try|pln|irr|dollar[s]?|euro[s]?|rial[s]?|lira|[€$£₺])?\s*` +
		`(?P<amount>\d+(?:[.,]\d+)?)` +
		`(?:\s*(?P<currency2>btc|eth|usd|eur|gbp|try|pln|irr|dollar[s]?|euro[s]?|rial[s]?|lira|[€$£₺]))?` +
		`)$`)

	matches := regex.FindStringSubmatch(input)
	if matches == nil {
		return 0, "", errors.New("unable to parse money string")
	}

	// Get named groups
	var rawAmount, rawCurrency string
	for i, name := range regex.SubexpNames() {
		switch name {
		case "amount":
			rawAmount = matches[i]
		case "currency1":
			if matches[i] != "" {
				rawCurrency = matches[i]
			}
		case "currency2":
			if matches[i] != "" {
				rawCurrency = matches[i]
			}
		}
	}

	if rawAmount == "" {
		return 0, "", errors.New("amount not found")
	}

	currency := NormalizeCurrency(rawCurrency)
	if currency == "" {
		return 0, "", errors.New("currency not recognized")
	}

	// Replace comma decimal separator to dot if needed
	rawAmount = strings.ReplaceAll(rawAmount, ",", ".")

	floatVal, err := strconv.ParseFloat(rawAmount, 64)
	if err != nil {
		return 0, "", err
	}

	mult := currencyMultiplier(currency)
	// round to avoid float precision errors
	amount := int64(floatVal*float64(mult) + 0.5)

	return amount, currency, nil
}

func formatWithCommas(n int64) string {
	s := strconv.FormatInt(n, 10)
	if len(s) <= 3 {
		return s
	}

	var b strings.Builder
	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}

	b.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}

func (m Money) Format(locale string) string {
	if !m.Amount.Valid || !m.Currency.Valid {
		return ""
	}

	if m.Currency.String == "BTC" {
		// BTC uses 8 decimals, stored as satoshis (int64)
		satoshis := m.Amount.Int64
		intPart := satoshis / 100_000_000
		fracPart := satoshis % 100_000_000

		// Format fractional part with leading zeros
		fracStr := fmt.Sprintf("%08d", fracPart)
		// Trim trailing zeros for nicer output (optional)
		fracStr = strings.TrimRight(fracStr, "0")

		if fracStr == "" {
			return fmt.Sprintf("%d BTC", intPart)
		}
		// Add thousands separators to intPart (optional)
		intStr := formatWithCommas(intPart)

		return fmt.Sprintf("%s.%s BTC", intStr, fracStr)
	}

	curCode := m.Currency.String
	moneyVal := money.New(m.Amount.Int64, curCode)
	return moneyVal.Display()
}

func NormalizeCurrency(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))

	currencyMap := map[string]string{
		// USD
		"$":          "USD",
		"usd":        "USD",
		"dollar":     "USD",
		"dollars":    "USD",
		"us dollar":  "USD",
		"us dollars": "USD",

		// EUR
		"€":     "EUR",
		"eur":   "EUR",
		"euro":  "EUR",
		"euros": "EUR",

		// PLN
		"pln":    "PLN",
		"zł":     "PLN",
		"zloty":  "PLN",
		"zlotys": "PLN",

		// IRR
		"irr":   "IRR",
		"rial":  "IRR",
		"rials": "IRR",

		// TRY
		"₺":     "TRY",
		"try":   "TRY",
		"lira":  "TRY",
		"liras": "TRY",

		// GBP
		"£":              "GBP",
		"gbp":            "GBP",
		"pound":          "GBP",
		"pounds":         "GBP",
		"british pound":  "GBP",
		"british pounds": "GBP",

		// JPY
		"¥":   "JPY",
		"jpy": "JPY",
		"yen": "JPY",

		// CHF (Swiss Franc)
		"chf":    "CHF",
		"franc":  "CHF",
		"francs": "CHF",

		// AUD
		"aud":               "AUD",
		"australian dollar": "AUD",

		// CAD
		"cad":             "CAD",
		"canadian dollar": "CAD",

		// BTC
		"btc": "BTC",

		// ETH
		"eth": "ETH",
	}

	if val, ok := currencyMap[input]; ok {
		return strings.ToUpper(val)
	}

	return ""
}

// Add sums two Money values; returns zero Money if currencies differ or invalid
func (m Money) Add(other Money) Money {
	if !m.Amount.Valid || !other.Amount.Valid || m.Currency.String != other.Currency.String {
		return Money{}
	}
	return NewMoney(m.Amount.Int64+other.Amount.Int64, m.Currency.String)
}

func (m Money) Plus(otherStr string) Money {
	other := NewMoneyAutoNull(otherStr)

	return m.Add(other)
}

// Sub subtracts other from m; returns zero Money if currencies differ or invalid
func (m Money) Sub(other Money) Money {
	if !m.Amount.Valid || !other.Amount.Valid || m.Currency.String != other.Currency.String {
		return Money{}
	}
	return NewMoney(m.Amount.Int64-other.Amount.Int64, m.Currency.String)
}

func (m Money) Minus(otherStr string) Money {
	other := NewMoneyAutoNull(otherStr)

	return m.Sub(other)
}

// Equal checks if two Money values are equal in amount and currency
func (m Money) Equal(other Money) bool {
	return m.Amount.Valid && other.Amount.Valid &&
		m.Amount.Int64 == other.Amount.Int64 &&
		m.Currency.String == other.Currency.String
}

func (m Money) IsSame(otherStr string) bool {
	other := NewMoneyAutoNull(otherStr)
	return m.Amount.Valid && other.Amount.Valid &&
		m.Amount.Int64 == other.Amount.Int64 &&
		m.Currency.String == other.Currency.String
}

// IsZero returns true if amount is zero or invalid
func (m Money) IsZero() bool {
	return !m.Amount.Valid || m.Amount.Int64 == 0
}

// Negate returns negative of Money value
func (m Money) Negate() Money {
	if !m.Amount.Valid {
		return Money{}
	}
	return NewMoney(-m.Amount.Int64, m.Currency.String)
}
