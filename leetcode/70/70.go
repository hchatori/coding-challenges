package main

import "fmt"

// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can
// you climb to the top?

// 1 <= n <= 45

// This solution uses tabulation

func climbStairs(n int) int {

	ways := make([]int, n+1)

	ways[0] = 1
	ways[1] = 1

	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]
}

func main() {
	n := 44
	fmt.Println(climbStairs(n))
}
