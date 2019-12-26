package dataconv

import "fmt"

// ShowConv demonstrates some type conversion
func ShowConv() {
	// int
	var a = 24

	// float
	var b = 2.0

	// convert the int to a float64 for this calculation
	c := float64(a) * b
	fmt.Println(c)

	// fmt.Sprintf is a good way to convert to strings
	precision := fmt.Sprintf("%.2f", b)

	// print the value and the types
	fmt.Printf("%s - %T\n", precision, precision)

}
