package main

import (
	"fmt"

	"github.com/mrojasb2000/go-programming-cookbook/chapter04/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
