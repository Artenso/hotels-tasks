package main

import (
	"math/big"
)

func findSimples(min, max int) []int {
	res := make([]int, 0, 4)

	for num := min; num <= max; num++ {
		if big.NewInt(int64(num)).ProbablyPrime(0) {
			res = append(res, num)
		}
	}
	return res
}
