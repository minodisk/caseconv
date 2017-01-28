// Package caseconv implements converting case functions.
package caseconv

import (
	"bytes"
	"strings"
	"unicode"
)

// LowerSnakeCase converts s to lower snake case.
func LowerSnakeCase(s string) string {
	return strings.ToLower(SnakeCase(s))
}

// UpperSnakeCase converts s to upper snake case.
func UpperSnakeCase(s string) string {
	return strings.ToUpper(SnakeCase(s))
}

// SnakeCase converts s to snake case.
func SnakeCase(s string) string {
	return Delimit(s, "_")
}

// Delimit slices s into all substrings at lower and upper
// boundaries and before numeric characters,
// and concatenates the substrings with sep.
func Delimit(s, sep string) string {
	var b bytes.Buffer
	for i, c := range s {
		switch {
		case upper(c):
			if i > 0 {
				p := rune(s[i-1])
				if lower(p) {
					b.WriteString(sep)
				}
			}
			b.WriteRune(c)
		case !(alphabetical(c) || numeric(c)):
			b.WriteString(sep)
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

// LowerCamelCase converts s to lower camel case.
func LowerCamelCase(s string) string {
	return CamelCase(s, unicode.ToLower)
}

// UpperCamelCase converts s to upper camel case.
func UpperCamelCase(s string) string {
	return CamelCase(s, unicode.ToUpper)
}

// ConvertFunc converts a given character to another case.
type ConvertFunc func(rune) rune

// CamelCase converts the characters after non alphabetical
// character and after numeric character in s with convert.
func CamelCase(s string, convert ConvertFunc) string {
	var b bytes.Buffer
	first := true
	apply := false
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		switch {
		case !(alphabetical(c) || numeric(c)):
			apply = true
		case first:
			first = false
			apply = false
			b.WriteRune(convert(c))
		case apply:
			apply = false
			b.WriteRune(unicode.ToUpper(c))
		default:
			apply = false
			b.WriteRune(c)
		}
	}
	return b.String()
}

// alphabetical reports whether c is a alphabet.
func alphabetical(c rune) bool {
	return lower(c) || upper(c)
}

// upper reports whether c is a upper case alphabet.
func upper(c rune) bool {
	return 'A' <= c && 'Z' >= c
}

// upper reports whether c is a lower case alphabet.
func lower(c rune) bool {
	return 'a' <= c && 'z' >= c
}

// upper reports whether c is a numeric character.
func numeric(c rune) bool {
	return '0' <= c && '9' >= c
}
