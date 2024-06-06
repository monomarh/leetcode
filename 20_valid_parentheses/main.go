package main

import "fmt"

func main() {
	fmt.Println(isValid("(([]){})"))
}

func isValid(s string) bool {
	var stack []int32
	for _, ch := range s {
		if len(stack) == 0 || ch-stack[len(stack)-1] != 2 && ch-stack[len(stack)-1] != 1 {
			stack = append(stack, ch)
		} else {
			stack = stack[:len(stack)-1]
		}

	}
	return len(stack) == 0
}
