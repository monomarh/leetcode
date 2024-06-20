// In the universe Earth C-137, Rick discovered a special form of magnetic force between two balls
// if they are put in his new invented basket. Rick has n empty baskets, the ith basket is at position[i],
// Morty has m balls and needs to distribute the balls into the baskets
// such that the minimum magnetic force between any two balls is maximum.
//
// Rick stated that magnetic force between two different balls at positions x and y is |x - y|.
//
// Given the integer array position and the integer m. Return the required force.
//
// Example 1:
//
// Input: position = [1,2,3,4,7], m = 3
// Output: 3
// Explanation: Distributing the 3 balls into baskets 1, 4 and 7
// will make the magnetic force between ball pairs [3, 3, 6].
// The minimum magnetic force is 3. We cannot achieve a larger minimum magnetic force than 3.
//
// Example 2:
//
// Input: position = [5,4,3,2,1,1000000000], m = 2
// Output: 999999999
// Explanation: We can use baskets 1 and 1000000000.
package main

import (
	"fmt"
	"sort"
)

var cases = []struct {
	position []int
	m        int
}{
	{
		position: []int{79, 74, 57, 22},
		m:        4,
	},
	{
		position: []int{1, 2, 3, 4, 7},
		m:        3,
	},
	{
		position: []int{5, 4, 3, 2, 1, 1000000000},
		m:        2,
	},
}

func main() {
	for i, c := range cases {
		result := maxDistance(c.position, c.m)

		fmt.Printf("Test case %d; input: (%#v, %d); output: %d\n", i, c.position, c.m, result)
	}
}

func maxDistance(position []int, m int) (force int) {
	sort.Ints(position)

	minForce, maxForce := 1, position[len(position)-1]-position[0]

	for minForce <= maxForce {
		mid := minForce + (maxForce-minForce)/2

		if canBePlaced(position, mid, m) {
			force = mid
			minForce = mid + 1
		} else {
			maxForce = mid - 1
		}
	}

	return
}

func canBePlaced(position []int, force int, balls int) bool {
	previous := position[0]
	count := 1

	for i := 1; i < len(position) && count < balls; i++ {
		if position[i]-previous >= force {
			count++
			previous = position[i]
		}
	}

	return count == balls
}
