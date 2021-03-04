package main

import "fmt"

func main() {
	// nums := []int{3, 0, 1}
	// nums := []int{0, 1}
	// nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	// nums := []int{0}
	// nums := []int{1}
	nums := []int{1, 2}
	fmt.Println(missingNumber(nums))
}

// Solution 3: Fast! Beats 88% of submissions.
func missingNumber(nums []int) int {

	// Use an equation to find the expected sum of numbers between 0 and n
	n := len(nums) + 1
	x := (n - 1) / 2
	y := (n - 1) % 2
	expectedSum := n*x + (n/2)*y

	// Find the actualSum of numbers in nums and search for a zero
	actualSum := 0
	foundZero := false

	for _, n := range nums {
		actualSum += n
		if n == 0 {
			foundZero = true
		}
	}

	if foundZero == false {
		return 0
	}

	// If there is a zero in nums and the difference between the expectedSum
	// and actualSum is zero, the missingNum is one greater than the largest num
	if expectedSum-actualSum == 0 {
		return len(nums) + 1
	}

	// Otherwise, the missingNum is the difference between expectedSum and
	// actualSum
	return expectedSum - actualSum
}

// Solution 2: Still slow. Creating a map and searching the map.
// type void struct{}

// func missingNumber(nums []int) int {
// 	foundNums := make(map[int]void)
// 	maxNumber := 0

// 	for _, n := range nums {
// 		foundNums[n] = void{}
// 		if n > maxNumber {
// 			maxNumber = n
// 		}
// 	}

// 	for i := 0; i <= maxNumber; i++ {
// 		if _, ok := foundNums[i]; !ok {
// 			return i
// 		}
// 	}

// 	return maxNumber + 1
// }

// Solution 1: Extremely slow solution. Sorting array.
// func missingNumber(nums []int) int {
// 	sort.Ints(nums)
// 	maxNumber := 0

// 	for _, n := range nums {
// 		if n > maxNumber {
// 			maxNumber = n
// 		}
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != i {
// 			return i
// 		}
// 	}

// 	return maxNumber + 1
// }
