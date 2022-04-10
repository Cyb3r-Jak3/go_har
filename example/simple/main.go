package main

import (
	"fmt"
	"log"

	"github.com/Cyb3r-Jak3/go_har"
)

func main() {
	har, err := har.ParseHar("testdata/Firefox.har")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(har.Version)
}
