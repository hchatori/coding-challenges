package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	triangle := [][]int{}

	for i := 0; i < numRows; i++ {
		triangleRow := []int{}

		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				triangleRow = append(triangleRow, 1)
			} else {
				if i >= 1 && j >= 1 {
					num1 := triangle[i-1][j-1]
					num2 := triangle[i-1][j]
					triangleRow = append(triangleRow, num1+num2)
				}
			}
		}
		triangle = append(triangle, triangleRow)
	}
	return triangle
}
