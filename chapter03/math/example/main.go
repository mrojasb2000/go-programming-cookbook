package main

import (
	"fmt"

	"github.com/mrojasb2000/go-programming-cookbook/chapter03/math"
)

func main() {
	math.Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", math.Fib(i))
	}
	fmt.Println()
}
