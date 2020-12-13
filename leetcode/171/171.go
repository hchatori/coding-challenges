package main

import (
	"fmt"
	"math"
)

func main() {
	s := "AAAA"
	fmt.Println(titleToNumber(s))
}

func titleToNumber(s string) int {
	r := []rune(s)

	placeValue := len(r) - 1
	total := 0
	for _, c := range r {
		total += int(math.Pow(26.0, float64(placeValue)) * float64(c-64))
		placeValue--
	}

	return total
}
