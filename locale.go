// Copyright (c) 2020 Bojan Zivanovic and contributors
// SPDX-License-Identifier: MIT

package currency

import "strings"

// Locale represents a Unicode locale identifier.
type Locale struct {
	Language string
	Script   string
	Region   string
}

// NewLocale creates a new Locale from its string representation.
func NewLocale(id string) Locale {
	// Normalize the ID ("SR_rs_LATN" => "sr-Latn-RS").
	id = strings.ToLower(id)
	id = strings.ReplaceAll(id, "_", "-")
	locale := Locale{}
	for i, part := range strings.Split(id, "-") {
		if i == 0 {
			locale.Language = part
			continue
		}
		partLen := len(part)
		if partLen == 4 {
			locale.Script = strings.Title(part)
			continue
		}
		if partLen == 2 || partLen == 3 {
			locale.Region = strings.ToUpper(part)
			continue
		}
	}

	return locale
}

// String returns the string representation of l.
func (l Locale) String() string {
	b := strings.Builder{}
	b.WriteString(l.Language)
	if l.Script != "" {
		b.WriteString("-")
		b.WriteString(l.Script)
	}
	if l.Region != "" {
		b.WriteString("-")
		b.WriteString(l.Region)
	}

	return b.String()
}

// IsEmpty returns whether l is empty.
func (l Locale) IsEmpty() bool {
	return l.Language == "" && l.Script == "" && l.Region == ""
}

// GetParent returns the parent locale for l.
//
//	Order:
// 	1. Language - Script - Region (e.g. "sr-Cyrl-RS")
// 	2. Language - Script (e.g. "sr-Cyrl")
// 	3. Language (e.g. "sr")
// 	4. English ("en")
// 	5. Empty locale ("")
//
// Note that according to CLDR rules, certain locales have special parents.
// For example, the parent for "es-AR" is "es-419", and for "sr-Latn" it is "en".
func (l Locale) GetParent() Locale {
	localeID := l.String()
	if localeID == "" || localeID == "en" {
		return Locale{}
	}
	if p, ok := parentLocales[localeID]; ok {
		return NewLocale(p)
	}

	if l.Region != "" {
		return Locale{Language: l.Language, Script: l.Script}
	} else if l.Script != "" {
		return Locale{Language: l.Language}
	} else {
		return Locale{Language: "en"}
	}
}
