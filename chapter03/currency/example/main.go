package main

import (
	"fmt"

	"github.com/mrojasb2000/go-programming-cookbook/chapter03/currency"
)

func main() {
	// start with our user input
	// of fiften dollars and 93 cents
	userInput := "15.93"

	pennies, err := currency.ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User input converted to %d pennies\n", pennies)

	// add 15 cents
	pennies += 15
	dollars := currency.ConvertPenniesToDollarString(pennies)

	fmt.Printf("Added 15 cents, new values is %s dollars\n", dollars)
}
