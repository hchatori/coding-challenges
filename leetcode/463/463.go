package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}

	fmt.Println(islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	maxCol := len(grid[0]) - 1
	maxRow := len(grid) - 1

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				continue
			}

			// for every water cell adjacent to a land cell, add 1 to perimeter
			if col-1 >= 0 && grid[row][col-1] == 0 || col-1 < 0 { // left
				perimeter++
			}
			if row-1 >= 0 && grid[row-1][col] == 0 || row-1 < 0 { // up
				perimeter++
			}
			if col+1 <= maxCol && grid[row][col+1] == 0 || col+1 > maxCol { // right
				perimeter++
			}
			if row+1 <= maxRow && grid[row+1][col] == 0 || row+1 > maxRow { // down
				perimeter++
			}
		}
	}

	return perimeter
}
