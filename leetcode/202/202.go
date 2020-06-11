package main

import (
	"fmt"
	"strconv"
	"strings"
)

var void struct{}

var sums map[int]struct{} = map[int]struct{}{}

func main() {
	fmt.Println(isHappy(13))
}

func isHappy(n int) bool {
	// create a slice of digits of n
	strOfn := strconv.Itoa(n)
	arrayOfStr := strings.Split(strOfn, "")
	digits := []int{}
	for _, s := range arrayOfStr {
		i, _ := strconv.Atoi(s)
		digits = append(digits, i)
	}

	// find sum of digits
	sum := 0
	for _, digit := range digits {
		sum += digit * digit
	}

	if _, ok := sums[sum]; ok {
		sums = map[int]struct{}{}
		return false
	}

	if sum == 1 {
		sums = map[int]struct{}{}
		return true
	}

	// store sum in a map
	sums[sum] = void

	return isHappy(sum)
}
