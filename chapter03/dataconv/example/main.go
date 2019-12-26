package main

import "github.com/mrojasb2000/go-programming-cookbook/chapter03/dataconv"

func main() {
	dataconv.ShowConv()
	if err := dataconv.Strconv(); err != nil {
		panic(err)
	}
	dataconv.Interfaces()
}
