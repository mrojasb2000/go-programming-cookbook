package main

import "github.com/mrojasb2000/go-programming-cookbook/chapter04/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
