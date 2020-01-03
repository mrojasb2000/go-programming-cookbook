package main

import (
	"fmt"

	"github.com/mrojasb2000/go-programming-cookbook/chapter04/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
