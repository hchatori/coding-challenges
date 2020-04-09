package main

import (
	"fmt"
)

func main() {
	fmt.Println(numFirstSundays())
}

func numFirstSundays() int {
	daysToAdvanceByMonth := [12]int{3, 0, 3, 2, 3, 2, 3, 3, 2, 3, 2, 3}
	dayOfWeek := 2 // 1 january 1901 is a tuesday
	numSundays := 0

	// yearNum = (1901 to 2000)
	for yearNum := 1; yearNum <= 100; yearNum++ {
		for i, daysToAdvance := range daysToAdvanceByMonth {
			if yearNum%4 == 0 && i == 1 { // if leap year and month is february
				dayOfWeek += (daysToAdvance + 1) % 7
			} else {
				dayOfWeek += daysToAdvance % 7
			}
			if dayOfWeek == 0 {
				numSundays++
			}
		}
	}

	return numSundays
}
