package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	// Create a map where they key is a number in arr and the value is
	// the frequency in which the number occurs in arr.

	counts := map[int]int{}
	for _, num := range arr {
		_, ok := counts[num]
		if !ok {
			counts[num] = 0;
		} else {
			counts[num] += 1
		}
	}

	// Create a map and check the map before adding
	freqs := map[int]bool{}
	for _, value := range counts {
		_, ok := freqs[value]
		if ok {
			return false;
		}
		freqs[value] = false;
	}

    return true;
}