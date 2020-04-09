package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Hello, how are you?"))
}

func reverseWords(s string) string {
	sepWords := strings.Fields(s)
	reversed := ""
	for index, word := range sepWords {
		letters := []rune(word)
		for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
			letters[i], letters[j] = letters[j], letters[i]
		}
		if index != 0 {
			reversed += " "
		}
		reversed += string(letters)
	}
	return reversed
}
