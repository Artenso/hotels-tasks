package main

import (
	"math/big"
)

func findSimples(min, max int) []int {
	// create results slice with len = 0 and
	// cap = 4 because when we get more elements in nums,
	// we will have fewer common divisors and
	// there is no point in reserving extra memory,
	// however, this value allows us to avoid two allocations
	// in cases where there are more than two divisors.
	res := make([]int, 0, 4)

	for num := min; num <= max; num++ {
		// this func divides num by the product of prime numbers
		// and then checks if the remainder is divisible by prime numbers
		if big.NewInt(int64(num)).ProbablyPrime(0) {
			res = append(res, num)
		}
	}
	return res
}
