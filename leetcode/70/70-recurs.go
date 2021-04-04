package main

import "fmt"

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}

	numWays := chooseSteps(n)

	return numWays
}

func chooseSteps(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	numWays := chooseSteps(n-1) + chooseSteps(n-2)

	return numWays
}

func main() {
	n := 44
	fmt.Println(climbStairs(n))
}
