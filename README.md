# go-caseconv [ ![Codeship Status for minodisk/go-caseconv](https://app.codeship.com/projects/900fad50-c6e0-0134-ad8d-6634c9578aaa/status?branch=master)](https://app.codeship.com/projects/198764) [![Go Report Card](https://goreportcard.com/badge/github.com/minodisk/go-caseconv)](https://goreportcard.com/report/github.com/minodisk/go-caseconv) [![GoDoc](https://godoc.org/github.com/minodisk/go-caseconv?status.png)](https://godoc.org/github.com/minodisk/go-caseconv)

Converting cases.

## Installation

```go
go get github.com/minodisk/go-caseconv
```

## Usage

```go
caseconv.LowerSnakeCase("Foo bar baz") // -> "foo_bar_baz"
caseconv.LowerSnakeCase("FooBarBaz0")  // -> "foo_bar_baz_0"

caseconv.UpperSnakeCase("Foo bar baz") // -> "FOO_BAR_BAZ"
caseconv.UpperSnakeCase("FooBarBaz0")  // -> "FOO_BAR_BAZ_0"

caseconv.LowerCamelCase("Foo bar baz") // -> "fooBarBaz"
caseconv.LowerCamelCase("foo_bar_baz") // -> "fooBarBaz"
caseconv.LowerCamelCase("FooBarBaz0")  // -> "fooBarBaz0"

caseconv.UpperCamelCase("Foo bar baz") // -> "FooBarBaz"
caseconv.UpperCamelCase("foo_bar_baz") // -> "FooBarBaz"
caseconv.UpperCamelCase("fooBarBaz0")  // -> "FooBarBaz0"
```
