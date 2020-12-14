package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		}

		if count > 0 && s[i] == ' ' {
			return count
		}

		if i == 0 && count == 0 {
			return 0
		}
	}

	return count
}
