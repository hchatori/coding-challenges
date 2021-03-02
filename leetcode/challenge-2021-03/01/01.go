package main

import (
	"fmt"
)

func main() {
	// candyType := []int{1, 1, 2, 2, 3, 3} // A: 3
	// candyType := []int{1, 1, 2, 3} // A: 2
	// candyType := []int{6, 6, 6, 6} // A: 1
	candyType := []int{0, 0, 0, 4} // A: 2

	fmt.Println(distributeCandies(candyType))
}

type void struct{}

func distributeCandies(candyType []int) int {
	maxAllowed := len(candyType) / 2
	c := make(map[int]void)

	for _, t := range candyType {
		c[t] = void{}
	}

	if len(c) < maxAllowed {
		return len(c)
	}

	return maxAllowed
}

// func distributeCandies(candyType []int) int {
// 	sort.Ints(candyType)
// 	maxAllowed := len(candyType) / 2
// 	currType := -1
// 	numTypes := 0

// 	for _, t := range candyType {
// 		if t != currType {
// 			currType = t
// 			numTypes++
// 		}
// 	}

// 	if numTypes < maxAllowed {
// 		return numTypes
// 	}

// 	return maxAllowed
// }
