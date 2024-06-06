package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeKdigits("1432219", 3))
}

func removeKdigits(num string, k int) string {
	stack := make([]rune, 0)

	for _, c := range num {
		for len(stack) > 0 && stack[len(stack)-1] > c && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}

		if len(stack) > 0 || c != '0' {
			stack = append(stack, c)
		}
	}

	for len(stack) > 0 && k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
