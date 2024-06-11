// A school is trying to take an annual photo of all the students.
//
//	The students are asked to stand in a single file line in non-decreasing order by height.
//	Let this ordering be represented by the integer array expected
//	where expected[i] is the expected height of the ith student in line.
//
// You are given an integer array heights representing the current order that the students are standing in.
//
//	Each heights[i] is the height of the ith student in line (0-indexed).
//
// Return the number of indices where heights[i] != expected[i].
//
// Example 1:
//
// Input: heights = [1,1,4,2,1,3]
// Output: 3
// Explanation:
// heights:  [1,1,4,2,1,3]
// expected: [1,1,1,2,3,4]
// Indices 2, 4, and 5 do not match.
//
// Example 2:
//
// Input: heights = [5,1,2,3,4]
// Output: 5
// Explanation:
// heights:  [5,1,2,3,4]
// expected: [1,2,3,4,5]
// All indices do not match.
//
// Example 3:
//
// Input: heights = [1,2,3,4,5]
// Output: 0
// Explanation:
// heights:  [1,2,3,4,5]
// expected: [1,2,3,4,5]
// All indices match.
package main

import (
	"fmt"
	"slices"
)

var cases = []struct {
	heights []int
}{
	{
		heights: []int{1, 1, 4, 2, 1, 3},
	},
	{
		heights: []int{5, 1, 2, 3, 4},
	},
	{
		heights: []int{1, 2, 3, 4, 5},
	},
	{
		heights: []int{2, 1, 2, 1, 1, 2, 2, 1},
	},
}

func main() {
	for i, c := range cases {
		result := heightChecker([]int{2, 1, 2, 1, 1, 2, 2, 1})

		fmt.Printf("Test case %d; input: (%#v); output: %d\n", i, c.heights, result)
	}
}

func heightChecker(heights []int) (mismatches int) {
	countSorted := make([]int, slices.Max(heights)+1)

	for _, val := range heights {
		countSorted[val]++
	}

	for i := 1; i < len(countSorted); i++ {
		countSorted[i] += countSorted[i-1]
	}

	for _, val := range heights {
		if countSorted[val]--; heights[countSorted[val]] != val {
			mismatches++
		}
	}

	return
}
