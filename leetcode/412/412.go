package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
	output := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		foo := ""

		if i%3 == 0 {
			foo += "Fizz"
		}
		if i%5 == 0 {
			foo += "Buzz"
		}
		if foo == "" {
			foo = strconv.Itoa(i)
		}

		output = append(output, foo)
	}

	return output
}
