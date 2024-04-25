package main

func findDevidors(nums []int) []int {
	devidor := 1
	// create results slice with len = 0 and
	// cap = 4 because when we get more elements in nums,
	// we will have fewer common divisors and
	// there is no point in reserving extra memory,
	// however, this value allows us to avoid two allocations
	// in cases where there are more than two divisors.
	res := make([]int, 0, 4)
	// while devidor < min num in nums
	// increase devidor and check remainder of the division
	// if remainder is not 0 we move to the next devidor
	// else we check other nums and add devidor to results
	for {
		devidor++
		isCommonDevidor := true
		for _, num := range nums {
			if devidor > num {
				return res
			}
			if num%devidor != 0 {
				isCommonDevidor = false
				break
			}
		}
		if isCommonDevidor {
			res = append(res, devidor)
		}
	}

}
