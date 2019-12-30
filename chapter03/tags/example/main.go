package main

import (
	"fmt"

	"github.com/mrojasb2000/go-programming-cookbook/chapter03/tags"
)

func main() {
	if err := tags.EmptyStruct(); err != nil {
		panic(err)
	}
	fmt.Println()
	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
