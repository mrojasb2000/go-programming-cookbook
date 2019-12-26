package math

import "math/big"

// global to memoize fib
var memoize map[int]*big.Int

func init() {
	// initialize the map
	memoize = make(map[int]*big.Int)
}

// Fib orints the nth digit of the fibonacci sequence
// it will return 1 for anything < 0 as well...
// it's calculated recursyvely and use big.Int since
// int64 will quickly overflow
func Fib(n int) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}

	// base case
	if n < 2 {
		return big.NewInt(1)
	}

	// initialize map then add previous 2 fib values
	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1))
	memoize[n].Add(memoize[n], Fib(n-2))

	// return result
	return memoize[n]
}
