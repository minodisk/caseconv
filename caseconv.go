// Package caseconv implements converting case functions.
package caseconv

import (
	"bytes"
	"strings"
	"unicode"
)

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
	skipped := false
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		switch {
		case !(alphabetical(c) || numeric(c)):
			skipped = true
		case first:
			b.WriteRune(convert(c))
			first = false
			skipped = false
		case skipped:
			b.WriteRune(unicode.ToUpper(c))
			skipped = false
		default:
			b.WriteRune(c)
			skipped = false
		}
	}
	return b.String()
}

// SnakeCase converts s to snake case.
func SnakeCase(s string) string {
	return Delimit(s, "_")
}

// LowerSnakeCase converts s to lower snake case.
func LowerSnakeCase(s string) string {
	return strings.ToLower(SnakeCase(s))
}

// UpperSnakeCase converts s to upper snake case.
func UpperSnakeCase(s string) string {
	return strings.ToUpper(SnakeCase(s))
}

// Hyphens converts s to hyphenated string.
func Hyphens(s string) string {
	return Delimit(s, "-")
}

// LowerHyphens converts s to lowercase hyphenated string.
func LowerHyphens(s string) string {
	return strings.ToLower(Hyphens(s))
}

// UpperHyphens converts s to uppercase hyphenated string.
func UpperHyphens(s string) string {
	return strings.ToUpper(Hyphens(s))
}

// Delimit slices s into all substrings at lower and upper
// boundaries and before numeric characters,
// and concatenates the substrings with sep.
func Delimit(s, sep string) string {
	var (
		b         bytes.Buffer
		separated bool
	)
	for i, c := range s {
		switch {
		case upper(c):
			if i > 0 {
				p := rune(s[i-1])
				if lower(p) {
					if !separated {
						b.WriteString(sep)
						separated = true
					}
				}
			}
			b.WriteRune(c)
			separated = false
		case !(alphabetical(c) || numeric(c)):
			if !separated {
				b.WriteString(sep)
				separated = true
			}
		default:
			b.WriteRune(c)
			separated = false
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
