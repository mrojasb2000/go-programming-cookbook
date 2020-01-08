package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	fmt.Println("Listening on port :8000")
	// we mount out single handler on port localhost:8000
	// to handle all request
	log.Panic(http.ListenAndServe("localhost:8000", http.HandlerFunc(wsHandler)))
}
