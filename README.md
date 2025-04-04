# GO_HAR

[![Go](https://github.com/Cyb3r-Jak3/go_har/actions/workflows/main.yml/badge.svg)](https://github.com/Cyb3r-Jak3/go_har/actions/workflows/main.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/go_har)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/go_har)  [![GoDoc](https://godoc.org/github.com/Cyb3r-Jak3/go_har?status.svg)](https://godoc.org/github.com/Cyb3r-Jak3/go_har)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Cyb3r-Jak3/go_har?style=flat-square) ![GitHub](https://img.shields.io/github/license/Cyb3r-Jak3/go_har?style=flat-square)
---

HAR is a GoLang module that parses HTTP Archive files and allows for easy access to pages and entries.

## Requirements

Go 1.16+

#### Support

This module is tested with GO versions 1.16, 1.17, 1.18. It is tested with HAR files from [Firefox](testdata/Firefox.har) and [Chrome](testdata/Chrome.har). PRs containing HAR files from other browsers are welcome.

## Getting started

#### Download

`go get -u github.com/Cyb3r-Jak3/go_har/v2`

#### Using

```golang
import "github.com/Cyb3r-Jak3/go_har/v2"

...
harFile, err := har.ParseHar("<filepath>")
if err != nil {
    log.Fatal(err)
}
```

[example](example/simple/main.go)
