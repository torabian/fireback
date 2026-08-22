package wallet

import (
	"fmt"
	"math/big"
)

// Money.go centralizes every bit of arithmetic this module does on wallet amounts, so
// there is exactly one place that ever touches the minor-units string representation.
// Every wallet/walletTransaction/walletPaymentAttempt "amount"-shaped field is a decimal
// string of integer minor-units (see Wallet.emi.yml's wallet.balance field doc) - e.g. for
// a currency with 2 decimals, "10050" means 100.50 units. This is deliberately never a
// float32/float64: float64 only carries ~15-17 significant decimal digits, which is fine
// for USD-scale amounts but silently lossy for 18-decimal crypto amounts, and money must
// never round in ways nobody asked for. math/big.Int has no such limit.

// ParseAmount parses a minor-units amount string into a *big.Int. Returns an error if s
// isn't a valid base-10 integer (including empty string) or is negative - every amount
// field in this module is a non-negative magnitude; direction/sign is carried separately
// (walletTransaction.direction, adjustBalance's direction field).
func ParseAmount(s string) (*big.Int, error) {
	if s == "" {
		return nil, fmt.Errorf("amount is empty")
	}
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil, fmt.Errorf("amount %q is not a valid integer", s)
	}
	if n.Sign() < 0 {
		return nil, fmt.Errorf("amount %q must not be negative", s)
	}
	return n, nil
}

// FormatAmount renders a *big.Int back into the minor-units decimal string stored on
// entities.
func FormatAmount(n *big.Int) string {
	return n.String()
}

// AddAmounts returns a+b as a minor-units string.
func AddAmounts(a, b string) (string, error) {
	an, err := ParseAmount(a)
	if err != nil {
		return "", err
	}
	bn, err := ParseAmount(b)
	if err != nil {
		return "", err
	}
	return FormatAmount(new(big.Int).Add(an, bn)), nil
}

// SubAmounts returns a-b as a minor-units string, and reports ok=false (no error - this
// is an expected, common outcome, not a failure) if the result would be negative, i.e. b
// exceeds a. Callers use this to detect "insufficient balance" without a separate
// comparison step.
func SubAmounts(a, b string) (result string, ok bool, err error) {
	an, err := ParseAmount(a)
	if err != nil {
		return "", false, err
	}
	bn, err := ParseAmount(b)
	if err != nil {
		return "", false, err
	}
	diff := new(big.Int).Sub(an, bn)
	if diff.Sign() < 0 {
		return "", false, nil
	}
	return FormatAmount(diff), true, nil
}

// IsZeroAmount reports whether the given minor-units string represents zero.
func IsZeroAmount(s string) (bool, error) {
	n, err := ParseAmount(s)
	if err != nil {
		return false, err
	}
	return n.Sign() == 0, nil
}
