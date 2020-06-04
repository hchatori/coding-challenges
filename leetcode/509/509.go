package main

import "fmt"

func main() {
	fmt.Println(fib(15))
}

var knownFib = map[int]int{}

func fib(N int) int {
	if val, ok := knownFib[N]; ok {
		return val
	}
	if N == 0 {
		knownFib[N] = 0
		return 0
	}
	if N == 1 {
		knownFib[N] = 1
		return 1
	}
	fibN := fib(N-2) + fib(N-1)
	knownFib[N] = fibN
	return fibN
}
