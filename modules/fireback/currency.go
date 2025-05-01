package fireback

import (
	"fmt"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// func CurrencyFormat(curr string, amount float64, lang language.Tag) string {
func CurrencyFormat(curr string, amount float64, query QueryDSL) string {
	if curr == "" || amount == 0 {
		return ""
	}

	lang := language.English
	if query.Language == "fa" {
		lang = language.Persian
	}
	cur, err := currency.ParseISO(curr)

	if err != nil {
		return fmt.Sprintf("%v", amount)
	}
	scale, _ := currency.Cash.Rounding(cur)
	dec := number.Decimal(amount, number.Scale(scale))
	p := message.NewPrinter(lang)
	return p.Sprintf("%v%v", dec, currency.Symbol(cur))
}
