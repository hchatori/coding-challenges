package main

import (
	"fmt"
)

func main() {
	// nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// nums := []int{1, 1}
	// nums := []int{1, 1, 2, 2}
	// nums := []int{2, 2}
	nums := []int{27, 40, 6, 21, 14, 36, 10, 19, 44, 10, 41, 26, 39, 20, 25, 19, 14, 7, 29, 27, 40, 38, 11, 44, 4, 6, 48, 39, 9, 13, 7, 45, 41, 23, 31, 8, 24, 1, 3, 5, 28, 11, 49, 29, 18, 4, 38, 32, 24, 15}
	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	existingNums := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if existingNums[nums[i]-1] == 1 {
			continue
		}

		existingNums[nums[i]-1] = 1
	}

	disappearedNums := []int{}

	for i := 0; i < len(existingNums); i++ {
		if existingNums[i] == 0 {
			disappearedNums = append(disappearedNums, i+1)
		}
	}

	return disappearedNums
}
