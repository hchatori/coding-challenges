package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./project-euler/22/names.txt")
	if err != nil {
		fmt.Println("file reading error", err)
		return
	}
	cleanedData := strings.ReplaceAll(string(data), "\"", "")

	namesList := strings.Split(cleanedData, ",")
	sort.Strings(namesList)

	total := 0
	for i, name := range namesList {
		total += nameScore(i+1, name)
	}

	fmt.Println(total)
}

func nameScore(count int, name string) int {
	nameScore := 0

	for _, letter := range name {
		nameScore += int(letter) - int('A') + 1
	}
	return nameScore * count
}
