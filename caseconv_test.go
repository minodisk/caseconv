package caseconv_test

import (
	"fmt"
	"testing"

	"github.com/minodisk/caseconv"
)

func TestLowerCamelCase(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "lower spaced string",
			Expected: "lowerSpacedString",
		},
		{
			Input:    "Upper Spaced String",
			Expected: "upperSpacedString",
		},
		{
			Input:    "lower_snake_case",
			Expected: "lowerSnakeCase",
		},
		{
			Input:    "Upper_Snake_Case",
			Expected: "upperSnakeCase",
		},
		{
			Input:    "lower-hyphens",
			Expected: "lowerHyphens",
		},
		{
			Input:    "Upper-Hyphens",
			Expected: "upperHyphens",
		},
		{
			Input:    "lowerCamelCase",
			Expected: "lowerCamelCase",
		},
		{
			Input:    "UpperCamelCase",
			Expected: "upperCamelCase",
		},
		{
			Input:    "numeric 00 string",
			Expected: "numeric00String",
		},
		{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "mixedVeryUgly00Case",
		},
		{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "obiWanBenKenobi",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			actual := caseconv.LowerCamelCase(c.Input)
			if actual != c.Expected {
				t.Errorf(`LowerCamelCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
			}
		})
	}
}

func TestUpperCamelCase(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "lower spaced string",
			Expected: "LowerSpacedString",
		},
		{
			Input:    "Upper Spaced String",
			Expected: "UpperSpacedString",
		},
		{
			Input:    "lower_snake_case",
			Expected: "LowerSnakeCase",
		},
		{
			Input:    "Upper_Snake_Case",
			Expected: "UpperSnakeCase",
		},
		{
			Input:    "lower-hyphens",
			Expected: "LowerHyphens",
		},
		{
			Input:    "Upper-Hyphens",
			Expected: "UpperHyphens",
		},
		{
			Input:    "lowerCamelCase",
			Expected: "LowerCamelCase",
		},
		{
			Input:    "UpperCamelCase",
			Expected: "UpperCamelCase",
		},
		{
			Input:    "numeric 00 string",
			Expected: "Numeric00String",
		},
		{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "MixedVeryUgly00Case",
		},
		{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "ObiWanBenKenobi",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			actual := caseconv.UpperCamelCase(c.Input)
			if actual != c.Expected {
				t.Errorf(`UpperCamelCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "lower spaced string",
			Expected: "lower_spaced_string",
		},
		{
			Input:    "Upper Spaced String",
			Expected: "Upper_Spaced_String",
		},
		{
			Input:    "lower_snake_case",
			Expected: "lower_snake_case",
		},
		{
			Input:    "Upper_Snake_Case",
			Expected: "Upper_Snake_Case",
		},
		{
			Input:    "lower-hyphens",
			Expected: "lower_hyphens",
		},
		{
			Input:    "Upper-Hyphens",
			Expected: "Upper_Hyphens",
		},
		{
			Input:    "lowerCamelCase",
			Expected: "lower_Camel_Case",
		},
		{
			Input:    "UpperCamelCase",
			Expected: "Upper_Camel_Case",
		},
		{
			Input:    "numeric 00 string",
			Expected: "numeric_00_string",
		},
		{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "_mixed_very_Ugly_00_Case_",
		},
		{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "Obi_Wan_Ben_Kenobi",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			actual := caseconv.SnakeCase(c.Input)
			if actual != c.Expected {
				t.Errorf(`LowerSnakeCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
			}
		})
	}
}

func TestLowerSnakeCase(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "lower spaced string",
			Expected: "lower_spaced_string",
		},
		{
			Input:    "Upper Spaced String",
			Expected: "upper_spaced_string",
		},
		{
			Input:    "lower_snake_case",
			Expected: "lower_snake_case",
		},
		{
			Input:    "Upper_Snake_Case",
			Expected: "upper_snake_case",
		},
		{
			Input:    "lower-hyphens",
			Expected: "lower_hyphens",
		},
		{
			Input:    "Upper-Hyphens",
			Expected: "upper_hyphens",
		},
		{
			Input:    "lowerCamelCase",
			Expected: "lower_camel_case",
		},
		{
			Input:    "UpperCamelCase",
			Expected: "upper_camel_case",
		},
		{
			Input:    "numeric 00 string",
			Expected: "numeric_00_string",
		},
		{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "_mixed_very_ugly_00_case_",
		},
		{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "obi_wan_ben_kenobi",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			actual := caseconv.LowerSnakeCase(c.Input)
			if actual != c.Expected {
				t.Errorf(`LowerSnakeCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
			}
		})
	}
}

func TestUpperSnakeCase(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "lower spaced string",
			Expected: "LOWER_SPACED_STRING",
		},
		{
			Input:    "Upper Spaced String",
			Expected: "UPPER_SPACED_STRING",
		},
		{
			Input:    "lower_snake_case",
			Expected: "LOWER_SNAKE_CASE",
		},
		{
			Input:    "Upper_Snake_Case",
			Expected: "UPPER_SNAKE_CASE",
		},
		{
			Input:    "lower-hyphens",
			Expected: "LOWER_HYPHENS",
		},
		{
			Input:    "Upper-Hyphens",
			Expected: "UPPER_HYPHENS",
		},
		{
			Input:    "lowerCamelCase",
			Expected: "LOWER_CAMEL_CASE",
		},
		{
			Input:    "UpperCamelCase",
			Expected: "UPPER_CAMEL_CASE",
		},
		{
			Input:    "numeric 00 string",
			Expected: "NUMERIC_00_STRING",
		},
		{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "_MIXED_VERY_UGLY_00_CASE_",
		},
		{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "OBI_WAN_BEN_KENOBI",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			actual := caseconv.UpperSnakeCase(c.Input)
			if actual != c.Expected {
				t.Errorf(`UpperSnakeCase("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
			}
		})
	}
}

func TestHyphens(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "lower spaced string",
			Expected: "lower-spaced-string",
		},
		{
			Input:    "Upper Spaced String",
			Expected: "Upper-Spaced-String",
		},
		{
			Input:    "lower_snake_case",
			Expected: "lower-snake-case",
		},
		{
			Input:    "Upper_Snake_Case",
			Expected: "Upper-Snake-Case",
		},
		{
			Input:    "lower-hyphens",
			Expected: "lower-hyphens",
		},
		{
			Input:    "Upper-Hyphens",
			Expected: "Upper-Hyphens",
		},
		{
			Input:    "lowerCamelCase",
			Expected: "lower-Camel-Case",
		},
		{
			Input:    "UpperCamelCase",
			Expected: "Upper-Camel-Case",
		},
		{
			Input:    "numeric 00 string",
			Expected: "numeric-00-string",
		},
		{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "-mixed-very-Ugly-00-Case-",
		},
		{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "Obi-Wan-Ben-Kenobi",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			actual := caseconv.Hyphens(c.Input)
			if actual != c.Expected {
				t.Errorf(`Hyphes("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
			}
		})
	}
}

func TestLowerHyphens(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "lower spaced string",
			Expected: "lower-spaced-string",
		},
		{
			Input:    "Upper Spaced String",
			Expected: "upper-spaced-string",
		},
		{
			Input:    "lower_snake_case",
			Expected: "lower-snake-case",
		},
		{
			Input:    "Upper_Snake_Case",
			Expected: "upper-snake-case",
		},
		{
			Input:    "lower-hyphens",
			Expected: "lower-hyphens",
		},
		{
			Input:    "Upper-Hyphens",
			Expected: "upper-hyphens",
		},
		{
			Input:    "lowerCamelCase",
			Expected: "lower-camel-case",
		},
		{
			Input:    "UpperCamelCase",
			Expected: "upper-camel-case",
		},
		{
			Input:    "numeric 00 string",
			Expected: "numeric-00-string",
		},
		{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "-mixed-very-ugly-00-case-",
		},
		{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "obi-wan-ben-kenobi",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			actual := caseconv.LowerHyphens(c.Input)
			if actual != c.Expected {
				t.Errorf(`LowerHyphes("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
			}
		})
	}
}
func TestUpperHyphens(t *testing.T) {
	t.Parallel()
	for i, c := range []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "lower spaced string",
			Expected: "LOWER-SPACED-STRING",
		},
		{
			Input:    "Upper Spaced String",
			Expected: "UPPER-SPACED-STRING",
		},
		{
			Input:    "lower_snake_case",
			Expected: "LOWER-SNAKE-CASE",
		},
		{
			Input:    "Upper_Snake_Case",
			Expected: "UPPER-SNAKE-CASE",
		},
		{
			Input:    "lower-hyphens",
			Expected: "LOWER-HYPHENS",
		},
		{
			Input:    "Upper-Hyphens",
			Expected: "UPPER-HYPHENS",
		},
		{
			Input:    "lowerCamelCase",
			Expected: "LOWER-CAMEL-CASE",
		},
		{
			Input:    "UpperCamelCase",
			Expected: "UPPER-CAMEL-CASE",
		},
		{
			Input:    "numeric 00 string",
			Expected: "NUMERIC-00-STRING",
		},
		{
			Input:    "_mixed veryUgly_00_Case_",
			Expected: "-MIXED-VERY-UGLY-00-CASE-",
		},
		{
			Input:    `Obi-Wan "Ben" Kenobi`,
			Expected: "OBI-WAN-BEN-KENOBI",
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			actual := caseconv.UpperHyphens(c.Input)
			if actual != c.Expected {
				t.Errorf(`UpperHyphens("%s") returns '%s', but expected '%s'`, c.Input, actual, c.Expected)
			}
		})
	}
}
