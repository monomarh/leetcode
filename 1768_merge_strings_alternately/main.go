package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mergeAlternately("abc", "zyx"))
}

func mergeAlternately(word1 string, word2 string) string {
	var newString strings.Builder
	maxLength := max(len(word1), len(word2))

	runes1, runes2 := []rune(word1), []rune(word2)

	for i := 0; i < maxLength; i++ {
		if i < len(word1) {
			newString.WriteRune(runes1[i])
		}
		if i < len(word2) {
			newString.WriteRune(runes2[i])
		}
	}

	return newString.String()
}
