package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	neg := false
	if x < 0 {
		x *= -1
		neg = true
	}

	// swap i and j chars in string
	runes := []rune(strconv.Itoa(x))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	rev, _ := strconv.Atoi(string(runes))

	if neg {
		rev *= -1
	}

	// if rev overflows a 32-bit integer, return 0
	if rev < math.MinInt32 || rev > math.MaxInt32 {
		return 0
	}

	return rev

}
