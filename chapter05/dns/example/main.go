package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mrojasb2000/go-programming-cookbook/chapter05/dns"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <address>\n", os.Args[0])
		os.Exit(1)
	}
	address := os.Args[1]
	lookup, err := dns.LookupAddress(address)
	if err != nil {
		log.Panicf("failed to llokup: %s", err.Error())
	}
	fmt.Println(lookup)
}
