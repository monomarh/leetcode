// Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.
//
// Example 1:
//
// Input: c = 5
// Output: true
// Explanation: 1 * 1 + 2 * 2 = 5
//
// Example 2:
//
// Input: c = 3
// Output: false
package main

import (
	"fmt"
	"math"
)

var cases = []struct {
	c int
}{
	{
		c: 5,
	},
	{
		c: 2,
	},
}

func main() {
	for i, c := range cases {
		result := judgeSquareSum(c.c)

		fmt.Printf("Test case %d; input: (%#v); output: %t\n", i, c.c, result)
	}
}

func judgeSquareSum(c int) bool {
	start, end := 0, int(math.Sqrt(float64(c)))

	for start <= end {
		val := start*start + end*end

		if val == c {
			return true
		} else if val > c {
			end--
		} else {
			start++
		}
	}

	return false
}
