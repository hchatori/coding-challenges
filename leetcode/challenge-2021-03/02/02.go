package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 2, 2, 4} //// A: {2,3}
	// nums := []int{2, 2} //// A: {2,1}
	// nums := []int{1, 3, 3} //// A: {3,2}
	// nums := []int{3, 2, 3, 4, 6, 5} //// A: {3, 1}
	nums := []int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9} //// A: {2, 10}
	fmt.Println(findErrorNums(nums))
}

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	errorNums := []int{0, 0}

	// Find duplicate
	for i, n := range nums {
		if i != len(nums)-1 && n == nums[i+1] {
			errorNums[0] = n
			break
		}
	}

	// Find missing number
	lastValidNum := 0
	for _, n := range nums {
		if n == lastValidNum+1 {
			lastValidNum++
		}
	}
	errorNums[1] = lastValidNum + 1

	return errorNums
}
