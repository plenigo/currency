// Copyright (c) 2020 Bojan Zivanovic and contributors
// SPDX-License-Identifier: MIT

package currency_test

import (
	"fmt"

	"github.com/plenigo/currency"
)

func ExampleNewAmount() {
	amount, _ := currency.NewAmount("24.49", "USD")
	fmt.Println(amount)
	fmt.Println(amount.Number())
	fmt.Println(amount.CurrencyCode())
	// Output: 24.49 USD
	// 24.49
	// USD
}

func ExampleAmount_ToMinorUnits() {
	firstAmount, _ := currency.NewAmount("20.99", "USD")
	secondAmount, _ := currency.NewAmount("50", "USD")
	fmt.Println(firstAmount.ToMinorUnits())
	fmt.Println(secondAmount.ToMinorUnits())
	// Output: 2099
	// 5000
}

func ExampleAmount_Convert() {
	amount, _ := currency.NewAmount("20.99", "USD")
	amount, _ = amount.Convert("EUR", "0.91")
	fmt.Println(amount)
	fmt.Println(amount.Round())
	// Output: 19.1009 EUR
	// 19.10 EUR
}

func ExampleAmount_Add() {
	firstAmount, _ := currency.NewAmount("20.99", "USD")
	secondAmount, _ := currency.NewAmount("3.50", "USD")
	totalAmount, _ := firstAmount.Add(secondAmount)
	fmt.Println(totalAmount)
	// Output: 24.49 USD
}

func ExampleAmount_Sub() {
	baseAmount, _ := currency.NewAmount("20.99", "USD")
	discountAmount, _ := currency.NewAmount("5.00", "USD")
	amount, _ := baseAmount.Sub(discountAmount)
	fmt.Println(amount)
	// Output: 15.99 USD
}

func ExampleAmount_Mul() {
	amount, _ := currency.NewAmount("20.99", "USD")
	taxAmount, _ := amount.Mul("0.20")
	fmt.Println(taxAmount)
	fmt.Println(taxAmount.Round())
	// Output: 4.1980 USD
	// 4.20 USD
}

func ExampleAmount_Div() {
	totalAmount, _ := currency.NewAmount("99.99", "USD")
	amount, _ := totalAmount.Div("3")
	fmt.Println(amount)
	// Output: 33.33 USD
}

func ExampleAmount_Round() {
	firstAmount, _ := currency.NewAmount("12.345", "USD")
	secondAmount, _ := currency.NewAmount("12.345", "JPY")
	fmt.Println(firstAmount.Round())
	fmt.Println(secondAmount.Round())
	// Output: 12.35 USD
	// 12 JPY
}

func ExampleAmount_RoundTo() {
	amount, _ := currency.NewAmount("12.345", "USD")
	for _, digits := range []uint8{4, 3, 2, 1, 0} {
		fmt.Println(amount.RoundTo(digits, currency.RoundHalfUp))
	}
	// Output: 12.3450 USD
	// 12.345 USD
	// 12.35 USD
	// 12.3 USD
	// 12 USD
}

func ExampleNewLocale() {
	firstLocale := currency.NewLocale("en-US")
	fmt.Println(firstLocale)
	fmt.Println(firstLocale.Language, firstLocale.Region)

	// Locale IDs are normalized.
	secondLocale := currency.NewLocale("sr_rs_latn")
	fmt.Println(secondLocale)
	fmt.Println(secondLocale.Language, secondLocale.Script, secondLocale.Region)
	// Output: en-US
	// en US
	// sr-Latn-RS
	// sr Latn RS
}

func ExampleLocale_GetParent() {
	locale := currency.NewLocale("sr-Cyrl-RS")
	for {
		fmt.Println(locale)
		locale = locale.GetParent()
		if locale.IsEmpty() {
			break
		}
	}
	// Output: sr-Cyrl-RS
	// sr-Cyrl
	// sr
	// en
}

func ExampleFormatter_Format() {
	locale := currency.NewLocale("tr")
	formatter := currency.NewFormatter(locale)
	amount, _ := currency.NewAmount("1245.988", "EUR")
	fmt.Println(formatter.Format(amount))

	formatter.MaxDigits = 2
	fmt.Println(formatter.Format(amount))

	formatter.NoGrouping = true
	amount, _ = currency.NewAmount("1245", "EUR")
	fmt.Println(formatter.Format(amount))

	formatter.MinDigits = 0
	fmt.Println(formatter.Format(amount))

	formatter.CurrencyDisplay = currency.DisplayNone
	fmt.Println(formatter.Format(amount))
	// Output: €1.245,988
	// €1.245,99
	// €1245,00
	// €1245
	// 1245
}

func ExampleFormatter_Parse() {
	locale := currency.NewLocale("tr")
	formatter := currency.NewFormatter(locale)

	amount, _ := formatter.Parse("€1.234,59", "EUR")
	fmt.Println(amount)

	amount, _ = formatter.Parse("EUR 1.234,59", "EUR")
	fmt.Println(amount)

	amount, _ = formatter.Parse("1.234,59", "EUR")
	fmt.Println(amount)
	// Output: 1234.59 EUR
	// 1234.59 EUR
	// 1234.59 EUR
}

func ExampleGetNumericCode() {
	numericCode, ok := currency.GetNumericCode("USD")
	fmt.Println(numericCode, ok)

	// Non-existent currency code.
	numericCode, ok = currency.GetNumericCode("XXX")
	fmt.Println(numericCode, ok)
	// Output: 840 true
	// 000 false
}

func ExampleGetDigits() {
	digits, ok := currency.GetDigits("USD")
	fmt.Println(digits, ok)

	// Non-existent currency code.
	digits, ok = currency.GetDigits("XXX")
	fmt.Println(digits, ok)
	// Output: 2 true
	// 0 false
}

func ExampleGetSymbol() {
	locale := currency.NewLocale("en")
	symbol, ok := currency.GetSymbol("USD", locale)
	fmt.Println(symbol, ok)

	// Non-existent currency code.
	symbol, ok = currency.GetSymbol("XXX", locale)
	fmt.Println(symbol, ok)
	// Output: $ true
	// XXX false
}
