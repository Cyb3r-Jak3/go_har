package main

import (
	"fmt"
	"log"

	"github.com/Cyb3r-Jak3/go_har/v2"
)

func main() {
	// Run from project root. To run from folder replace with "../../testdata/Firefox.har"
	har, err := har.ParseHar("testdata/Firefox.har")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(har.Version)
}
