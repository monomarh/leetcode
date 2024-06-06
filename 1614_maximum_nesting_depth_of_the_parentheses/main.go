package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}

func maxDepth(s string) (maxDepth int) {
	opened := 0

	for _, val := range s {
		if val == '(' {
			opened++
			maxDepth = max(maxDepth, opened)
		}
		if val == ')' {
			opened--
		}
	}

	return
}
