// Copyright (c) 2020 twihike. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// Package strcase is a string case converter between camel case, snake case
// and kebab case.
package strcase

import (
	"strings"
)

// ToLowerSnake converts to a lower case string with underscores between words.
func ToLowerSnake(s string) string {
	return addDelimiter(s, '_', false)
}

// ToUpperSnake converts to an upper case string with underscores between words.
func ToUpperSnake(s string) string {
	return addDelimiter(s, '_', true)
}

// ToLowerKebab converts to a lower case string with hyphens between words.
func ToLowerKebab(s string) string {
	return addDelimiter(s, '-', false)
}

// ToUpperKebab converts to an upper case string with hyphens between words.
func ToUpperKebab(s string) string {
	return addDelimiter(s, '-', true)
}

func addDelimiter(s string, d rune, toUpper bool) string {
	var sb strings.Builder
	runes := []rune(s)
	prevKind := kindUnknown

	for i, r := range runes {
		k := kind(r)
		switch k {
		case kindUpper:
			if prevKind == kindUpper {
				if i < len(runes)-1 {
					nextKind := kind(runes[i+1])
					if nextKind != kindUpper {
						sb.WriteRune(d)
					}
				}
			} else if prevKind != kindUnknown && prevKind != kindDelimiter {
				sb.WriteRune(d)
			}
			if toUpper {
				sb.WriteRune(r)
			} else {
				sb.WriteRune(r + 'a' - 'A')
			}
		case kindLower:
			if toUpper {
				sb.WriteRune(r + 'A' - 'a')
			} else {
				sb.WriteRune(r)
			}
		case kindNumeric:
			sb.WriteRune(r)
		case kindDelimiter:
			sb.WriteRune(d)
		case kindOther:
			sb.WriteRune(r)
		}
		prevKind = k
	}
	return sb.String()
}

// ToLowerCamel converts to capitalized words.
// However, the first word is lower case.
// It also removes the delimiter between words.
// If s is already camel case, it is not supported.
func ToLowerCamel(s string) string {
	return otherToCamel(s, false)
}

// ToUpperCamel converts to capitalized words.
// It also removes the delimiter between words.
// If s is already camel case, it is not supported.
func ToUpperCamel(s string) string {
	return otherToCamel(s, true)
}

func otherToCamel(s string, toUpper bool) string {
	var sb strings.Builder
	isCap := toUpper
	first := true

	for _, r := range s {
		k := kind(r)
		switch k {
		case kindUpper:
			if isCap {
				sb.WriteRune(r)
			} else {
				sb.WriteRune(r + 'a' - 'A')
			}
		case kindLower:
			if isCap {
				sb.WriteRune(r + 'A' - 'a')
			} else {
				sb.WriteRune(r)
			}
		case kindNumeric, kindOther:
			sb.WriteRune(r)
		}

		if k == kindDelimiter {
			if !first {
				isCap = true
			}
		} else {
			if first {
				first = false
			}
			isCap = false
		}
	}

	return sb.String()
}

// UpperCamelToLowerCamel converts an upper camel case string
// into a lower camel case string.
func UpperCamelToLowerCamel(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for i, r := range s {
		if i == 0 {
			if k := kind(r); k == kindUpper {
				sb.WriteRune(r + 'a' - 'A')
				continue
			}
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

// LowerCamelToUpperCamel converts a lower camel case string
// into an upper camel case string.
func LowerCamelToUpperCamel(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for i, r := range s {
		if i == 0 {
			if k := kind(r); k == kindLower {
				sb.WriteRune(r + 'A' - 'a')
				continue
			}
		}
		sb.WriteRune(r)
	}
	return sb.String()
}
