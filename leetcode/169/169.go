package main

import "fmt"

func main() {
	// nums := []int{2, 2, 1, 1, 1, 2, 2}
	nums := []int{1}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	maxCount := 0
	maxVal := 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num]++
		}

		if m[num] > maxCount {
			maxCount = m[num]
			maxVal = num

			if maxCount > len(nums)/2 {
				return maxVal
			}
		}
	}

	return maxVal
}
