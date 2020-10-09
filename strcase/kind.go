// Copyright (c) 2020 twihike. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package strcase

const (
	kindLower = iota
	kindUpper
	kindNumeric
	kindDelimiter
	kindOther
	kindUnknown
)

func kind(r rune) int {
	if isLower(r) {
		return kindLower
	}
	if isUpper(r) {
		return kindUpper
	}
	if isNumeric(r) {
		return kindNumeric
	}
	if isDelimiter(r) {
		return kindDelimiter
	}
	return kindOther
}

func isLower(r rune) bool {
	return 'a' <= r && r <= 'z'
}

func isUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func isNumeric(r rune) bool {
	return '0' <= r && r <= '9'
}

func isDelimiter(r rune) bool {
	return r == ' ' || r == '_' || r == '-'
}
