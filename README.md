# HAR
![Go](https://github.com/Cyb3r-Jak3/go_har/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Cyb3r-Jak3/go_har)](https://goreportcard.com/report/github.com/Cyb3r-Jak3/go_har)  [![GoDoc](https://godoc.org/github.com/Cyb3r-Jak3/go_har?status.svg)](https://godoc.org/github.com/Cyb3r-Jak3/go_har)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Cyb3r-Jak3/go_har?style=flat-square) ![GitHub](https://img.shields.io/github/license/Cyb3r-Jak3/go_har?style=flat-square)
---

HAR is a GoLang module that parses HTTP Archive files and allows for easy access to pages and entries.

## Requirements

Go 1.11

## Getting started

**Download**  
`go get github.com/Cyb3r-Jak3/go_har/`

**Use**
```golang
import "github.com/Cyb3r-Jak3/go_har"

...
har, err := go_har.parseHar("<filepath>")
if err != nil {
    log.Fatal(err)
}
```