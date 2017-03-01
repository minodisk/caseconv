# caseconv [ ![Codeship Status for minodisk/caseconv](https://img.shields.io/codeship/900fad50-c6e0-0134-ad8d-6634c9578aaa/master.svg?style=flat)](https://app.codeship.com/projects/198764) [![Go Report Card](https://goreportcard.com/badge/github.com/minodisk/caseconv)](https://goreportcard.com/report/github.com/minodisk/caseconv) [![Coverage Status](https://img.shields.io/coveralls/minodisk/caseconv/master.svg?style=flat)](https://coveralls.io/github/minodisk/caseconv?branch=master) [![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat)](https://godoc.org/github.com/minodisk/caseconv) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Converting cases.

## Installation

```bash
go get github.com/minodisk/caseconv
```

## Usage

```go
caseconv.LowerCamelCase(`Obi-Wan "Ben" Kenobi`) // -> "obiWanBenKenobi"
caseconv.UpperCamelCase(`Obi-Wan "Ben" Kenobi`) // -> "ObiWanBenKenobi"
caseconv.SnakeCase(`Obi-Wan "Ben" Kenobi`)      // -> "Obi_Wan_Ben_Kenobi"
caseconv.LowerSnakeCase(`Obi-Wan "Ben" Kenobi`) // -> "obi_wan_ben_kenobi"
caseconv.UpperSnakeCase(`Obi-Wan "Ben" Kenobi`) // -> "OBI_WAN_BEN_KENOBI"
caseconv.Hyphens(`Obi-Wan "Ben" Kenobi`)        // -> "Obi-Wan-Ben-Kenobi"
caseconv.LowerHyphens(`Obi-Wan "Ben" Kenobi`)   // -> "obi-wan-ben-kenobi"
caseconv.UpperHyphens(`Obi-Wan "Ben" Kenobi`)   // -> "OBI-WAN-BEN-KENOBI"
```
