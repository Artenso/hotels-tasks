package main

func findDevidors(nums []int) []int {
	devidor := 1
	res := make([]int, 0, 4)
	for {
		devidor++
		isCommonDevidor := true
		for _, num := range nums {
			if devidor > num {
				return res
			}
			if num%devidor != 0 {
				isCommonDevidor = false
			}
		}
		if isCommonDevidor {
			res = append(res, devidor)
		}
	}

}
