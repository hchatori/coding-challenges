package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}

func findDuplicates(nums []int) []int {
	numCount := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numCount[nums[i]-1] = numCount[nums[i]-1] + 1
	}

	duplicates := []int{}
	for i := 0; i < len(numCount); i++ {
		if numCount[i] == 2 {
			duplicates = append(duplicates, i+1)
		}
	}
	return duplicates
}
