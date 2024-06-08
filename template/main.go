// Problem description
package main

import "fmt"

var cases = []struct {
	input []int
}{
	{
		input: []int{1, 2, 3},
	},
}

func main() {
	for i, c := range cases {
		result := exampleFunc(c.input)

		fmt.Printf("Test case %d; input: (%#v); output: %t\n", i, c.input, result)
	}
}

func exampleFunc(nums []int) bool {
	if len(nums) > 0 {
		return false
	}
	return true
}
