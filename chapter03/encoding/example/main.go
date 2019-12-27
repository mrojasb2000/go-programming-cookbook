package main

import "github.com/mrojasb2000/go-programming-cookbook/chapter03/encoding"

func main() {
	if err := encoding.Base64Example(); err != nil {
		panic(err)
	}

	if err := encoding.Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := encoding.GoExample(); err != nil {
		panic(err)
	}
}
