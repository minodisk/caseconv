# go-caseconv [ ![Codeship Status for minodisk/go-caseconv](https://app.codeship.com/projects/900fad50-c6e0-0134-ad8d-6634c9578aaa/status?branch=master)](https://app.codeship.com/projects/198764) [![Go Report Card](https://goreportcard.com/badge/github.com/minodisk/go-caseconv)](https://goreportcard.com/report/github.com/minodisk/go-caseconv) [![GoDoc](https://godoc.org/github.com/minodisk/go-caseconv?status.png)](https://godoc.org/github.com/minodisk/go-caseconv)

Converting cases.

## Installation

```bash
go get github.com/minodisk/go-caseconv
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
