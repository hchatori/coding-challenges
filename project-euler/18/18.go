package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maxPathSum())
}

func maxPathSum() int {
	triangle := [][]int{}

	file, err := os.Open("project-euler/18/triangle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		triangleRowString := strings.Fields(scanner.Text())
		triangleRow := []int{}
		for _, s := range triangleRowString {
			n, _ := strconv.Atoi(s) // guaranteed s is a number, ignore err
			triangleRow = append(triangleRow, n)
		}
		triangle = append(triangle, triangleRow)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for row := len(triangle) - 1; row > 0; row-- {
		for col := 0; col < len(triangle[row])-1; col++ {
			first := triangle[row][col]
			second := triangle[row][col+1]
			if first >= second {
				triangle[row-1][col] += first
			} else {
				triangle[row-1][col] += second
			}
		}
	}

	return triangle[0][0]
}
