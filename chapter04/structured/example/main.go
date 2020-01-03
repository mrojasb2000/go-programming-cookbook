package main

import "fmt"

import "github.com/mrojasb2000/go-programming-cookbook/chapter04/structured"

func main() {
	fmt.Println("Logrus:")
	structured.Logrus()

	fmt.Println()
	fmt.Println("Apex:")
	structured.Apex()
}

/*
Output:
Logrus:
WARN[0000] warning!                                      complex_struct="{Something happened Just now}" id=123 success=true
ERRO[0000] Error!                                        complex_struct="{Something happened Just now}" id=123 success=true

Apex:
  INFO[0000] ThrowError                id=123
 ERROR[0000] ThrowError                duration=208ns error=a crazy failure id=123
 ERROR[0000] an error occurred         error=a crazy failure
*/
