package main

import "github.com/mrojasb2000/go-programming-cookbook/chapter01/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}

}
