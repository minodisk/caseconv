package caseconv_test

import (
	"testing"

	caseconv "github.com/minodisk/go-caseconv"
)

func TestLowerCamelCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "lowerSpacedString",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "upperSpacedString",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "lowerSnakeCase",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "upperSnakeCase",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "lowerHyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "upperHyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "lowerCamelCase",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "upperCamelCase",
		},
		Case{
			Input:    "numeric 00 string",
			Expected: "numeric00String",
		},
		Case{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "mixedVeryUgly00Case",
		},
		Case{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "obiWanBenKenobi",
		},
	}
	for _, c := range cases {
		actual := caseconv.LowerCamelCase(c.Input)
		if actual != c.Expected {
			t.Errorf(`LowerCamelCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}

func TestUpperCamelCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "LowerSpacedString",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "UpperSpacedString",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "LowerSnakeCase",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "UpperSnakeCase",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "LowerHyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "UpperHyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "LowerCamelCase",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "UpperCamelCase",
		},
		Case{
			Input:    "numeric 00 string",
			Expected: "Numeric00String",
		},
		Case{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "MixedVeryUgly00Case",
		},
		Case{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "ObiWanBenKenobi",
		},
	}
	for _, c := range cases {
		actual := caseconv.UpperCamelCase(c.Input)
		if actual != c.Expected {
			t.Errorf(`UpperCamelCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}

func TestSnakeCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "lower_spaced_string",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "Upper_Spaced_String",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "lower_snake_case",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "Upper_Snake_Case",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "lower_hyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "Upper_Hyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "lower_Camel_Case",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "Upper_Camel_Case",
		},
		Case{
			Input:    "numeric 00 string",
			Expected: "numeric_00_string",
		},
		Case{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "_mixed_very_Ugly_00_Case_",
		},
		Case{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "Obi_Wan_Ben_Kenobi",
		},
	}
	for _, c := range cases {
		actual := caseconv.SnakeCase(c.Input)
		if actual != c.Expected {
			t.Errorf(`LowerSnakeCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}

func TestLowerSnakeCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "lower_spaced_string",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "upper_spaced_string",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "lower_snake_case",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "upper_snake_case",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "lower_hyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "upper_hyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "lower_camel_case",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "upper_camel_case",
		},
		Case{
			Input:    "numeric 00 string",
			Expected: "numeric_00_string",
		},
		Case{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "_mixed_very_ugly_00_case_",
		},
		Case{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "obi_wan_ben_kenobi",
		},
	}
	for _, c := range cases {
		actual := caseconv.LowerSnakeCase(c.Input)
		if actual != c.Expected {
			t.Errorf(`LowerSnakeCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}

func TestUpperSnakeCase(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "LOWER_SPACED_STRING",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "UPPER_SPACED_STRING",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "LOWER_SNAKE_CASE",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "UPPER_SNAKE_CASE",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "LOWER_HYPHENS",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "UPPER_HYPHENS",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "LOWER_CAMEL_CASE",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "UPPER_CAMEL_CASE",
		},
		Case{
			Input:    "numeric 00 string",
			Expected: "NUMERIC_00_STRING",
		},
		Case{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "_MIXED_VERY_UGLY_00_CASE_",
		},
		Case{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "OBI_WAN_BEN_KENOBI",
		},
	}
	for _, c := range cases {
		actual := caseconv.UpperSnakeCase(c.Input)
		if actual != c.Expected {
			t.Errorf(`UpperSnakeCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}

func TestHyphens(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "lower-spaced-string",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "Upper-Spaced-String",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "lower-snake-case",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "Upper-Snake-Case",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "lower-hyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "Upper-Hyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "lower-Camel-Case",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "Upper-Camel-Case",
		},
		Case{
			Input:    "numeric 00 string",
			Expected: "numeric-00-string",
		},
		Case{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "-mixed-very-Ugly-00-Case-",
		},
		Case{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "Obi-Wan-Ben-Kenobi",
		},
	}
	for _, c := range cases {
		actual := caseconv.Hyphens(c.Input)
		if actual != c.Expected {
			t.Errorf(`Hyphes("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}

func TestLowerHyphens(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "lower-spaced-string",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "upper-spaced-string",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "lower-snake-case",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "upper-snake-case",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "lower-hyphens",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "upper-hyphens",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "lower-camel-case",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "upper-camel-case",
		},
		Case{
			Input:    "numeric 00 string",
			Expected: "numeric-00-string",
		},
		Case{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "-mixed-very-ugly-00-case-",
		},
		Case{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "obi-wan-ben-kenobi",
		},
	}
	for _, c := range cases {
		actual := caseconv.LowerHyphens(c.Input)
		if actual != c.Expected {
			t.Errorf(`LowerHyphes("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}
func TestUpperHyphens(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "lower spaced string",
			Expected: "LOWER-SPACED-STRING",
		},
		Case{
			Input:    "Upper Spaced String",
			Expected: "UPPER-SPACED-STRING",
		},
		Case{
			Input:    "lower_snake_case",
			Expected: "LOWER-SNAKE-CASE",
		},
		Case{
			Input:    "Upper_Snake_Case",
			Expected: "UPPER-SNAKE-CASE",
		},
		Case{
			Input:    "lower-hyphens",
			Expected: "LOWER-HYPHENS",
		},
		Case{
			Input:    "Upper-Hyphens",
			Expected: "UPPER-HYPHENS",
		},
		Case{
			Input:    "lowerCamelCase",
			Expected: "LOWER-CAMEL-CASE",
		},
		Case{
			Input:    "UpperCamelCase",
			Expected: "UPPER-CAMEL-CASE",
		},
		Case{
			Input:    "numeric 00 string",
			Expected: "NUMERIC-00-STRING",
		},
		Case{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "-MIXED-VERY-UGLY-00-CASE-",
		},
		Case{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "OBI-WAN-BEN-KENOBI",
		},
	}
	for _, c := range cases {
		actual := caseconv.UpperHyphens(c.Input)
		if actual != c.Expected {
			t.Errorf(`UpperHyphens("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}
