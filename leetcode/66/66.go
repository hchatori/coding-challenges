package main

import "fmt"

func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if carry+digits[i] != 10 {
			digits[i] += carry
			return digits
		}
		digits[i] = 0
		if i-1 == -1 {
			digits = append([]int{carry}, digits...)
		}
	}
	return digits
}
