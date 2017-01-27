package caseconv_test

import (
	"testing"

	caseconv "github.com/minodisk/go-caseconv"
)

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
	}
	for _, c := range cases {
		actual := caseconv.UpperSnakeCase(c.Input)
		if actual != c.Expected {
			t.Errorf(`UpperSnakeCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}

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
	}
	for _, c := range cases {
		actual := caseconv.UpperCamelCase(c.Input)
		if actual != c.Expected {
			t.Errorf(`UpperCamelCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
		}
	}
}
