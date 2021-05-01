package main

import "fmt"

func candy(ratings []int) int {
	candies := (make([]int, len(ratings)))
	for i := 0; i < len(ratings); i++ {
		candies[i] = 1
	}

	didUpdate := true
	for didUpdate {
		didUpdate = false
		for i := 0; i < len(ratings)-1; i++ {
			// If curr is less than next
			if ratings[i] < ratings[i+1] && candies[i] >= candies[i+1] {
				candies[i+1] = candies[i] + 1
				didUpdate = true
			}

			// If curr is greater than next
			if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
				didUpdate = true
			}

			if i > 0 {
				// If curr is less than prev
				if ratings[i] < ratings[i-1] && candies[i] >= candies[i-1] {
					candies[i-1] = candies[i] + 1
					didUpdate = true
				}
			}
		}
	}

	numCandies := 0
	for i := 0; i < len(candies); i++ {
		numCandies += candies[i]
	}

	// for i := 0; i < len(candies); i++ {
	// 	fmt.Println(candies[i])
	// }

	return numCandies
}

func main() {
	// ratings := []int{1, 2, 3} // 6
	// ratings := []int{1, 0, 2} // 5
	// ratings := []int{1, 2, 2} // 4
	// ratings := []int{1, 2, 87, 87, 87, 2, 1} // 13
	// ratings := []int{1, 6, 10, 8, 7, 3, 2} // 18
	// ratings := []int{5, 3, 4, 6, 6, 4, 3, 5, 1} // 17
	// ratings := []int{1, 3, 2, 2, 1} // 7
	ratings := []int{1, 2, 3, 5, 4, 3, 2, 1, 4, 3, 2, 1} // 31
	fmt.Println(candy(ratings))
}
