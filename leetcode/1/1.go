package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	targetNums := map[int]int{}
	for i := 0; i < len(nums); i++ {
		targetNums[target-nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if j, ok := targetNums[nums[i]]; ok {
			if i != j {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
