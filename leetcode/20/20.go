package main

import (
	"fmt"
)

func main() {
	// s := "([)]" // false
	s := "()}" // false
	// s := "{[]}" // true
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		// open char
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}

		// close char

		// unexpected: stack already empty
		if len(stack) == 0 {
			return false
		}

		// unexpected: close char does not match with top of stack
		topC := stack[len(stack)-1]
		if (c == ')' && topC != '(') || (c == ']' && topC != '[') || (c == '}' && topC != '{') {
			return false
		}

		// found match - pop from stack
		stack = stack[:len(stack)-1]
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
