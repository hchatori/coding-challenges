package main

import "fmt"

func main() {
	input := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(input, 9))
}

func search(nums []int, target int) int {
	loIndex := 0
	midIndex := (len(nums) - 1) / 2
	hiIndex := len(nums) - 1
	for {
		mid := (len(nums) - 1) / 2
		if target == nums[0] {
			return loIndex
		}
		if target == nums[mid] {
			return midIndex
		}
		if target == nums[len(nums)-1] {
			return hiIndex
		}
		if len(nums) < 4 {
			return -1
		}
		if target < nums[mid] {
			nums = nums[0 : mid+1]
			hiIndex = midIndex
			midIndex = (hiIndex + loIndex) / 2
			continue
		}
		if target > nums[mid] {
			nums = nums[mid+1 : len(nums)]
			loIndex = midIndex + 1
			midIndex = (hiIndex + loIndex) / 2
			continue
		}
	}
}
