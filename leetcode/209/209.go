package main

import (
	"fmt"
)

func main() {
	// s := 7
	// nums := []int{2, 3, 1, 2, 4, 3}

	// s := 11
	// nums := []int{1, 2, 3, 4, 5}

	s := 213
	nums := []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	fmt.Println(minSubArrayLen(s, nums))
}

func minSubArrayLen(s int, nums []int) int {
	sum := 0
	counter := 0
	fewestNums := s

	for j := 0; j <= len(nums)-1; j++ {

		sum = nums[j]
		counter++

		if sum >= s {
			return 1
		}

		for k := j + 1; k <= len(nums)-1; k++ {
			sum += nums[k]
			counter++

			if sum >= s {
				if counter < fewestNums {
					fewestNums = counter
				}
				sum = 0
				counter = 0
				break
			}

			if k == len(nums)-1 {
				sum = 0
				counter = 0
			}
		}

	}

	if fewestNums == s {
		return 0
	}

	return fewestNums
}
