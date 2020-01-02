package main

import (
	"fmt"

	"github.com/mrojasb2000/go-programming-cookbook/chapter04/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
