package main

import "math"

func findDevidors(nums []int) []int {

	// if there is no values in nums
	if len(nums) < 1 {
		return []int{}
	}

	// if there is one value in nums
	if len(nums) == 1 {
		return getNumDevidors(nums[0])
	}

	// we get greatest common devidor for all nums
	// and then get its devisors, so they will be
	// common devisors
	gcd := getGCD(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		gcd = getGCD(gcd, nums[i])
	}
	return getNumDevidors(gcd)

}

// getNumDevidors gets all devisors of one num
func getNumDevidors(num int) []int {

	// create map to avoid duplicate values
	devs := make(map[int]struct{})

	// create results slice with len = 0 and
	// cap = 4 because when we get more elements in nums,
	// we will have fewer common divisors and
	// there is no point in reserving extra memory,
	// however, this value allows us to avoid two allocations
	// in cases where there are more than two divisors.
	res := make([]int, 0, 4)

	// we have to look for divisors up to the root of the num,
	// because then there will be paired divisors and we get them like
	// num/dev
	for i := 1; i <= int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			devs[i] = struct{}{}
			devs[num/i] = struct{}{}
		}
	}

	// delete 1, because there was no 1 in the example
	delete(devs, 1)

	// put devisors in slice
	for dev, _ := range devs {
		res = append(res, dev)
	}
	return res
}

// getGCD gets greatest common devidor
// via Euclidean algorithm
func getGCD(num1, num2 int) int {
	for num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}
	return num1
}
