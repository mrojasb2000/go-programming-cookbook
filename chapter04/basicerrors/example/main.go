package main

import (
	"fmt"

	"github.com/mrojasb2000/go-programming-cookbook/chapter04/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
