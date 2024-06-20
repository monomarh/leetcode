// You have n jobs and m workers. You are given three arrays: difficulty, profit, and worker where:
//
// difficulty[i] and profit[i] are the difficulty and the profit of the ith job, and
// worker[j] is the ability of jth worker (i.e., the jth worker can only complete a job with difficulty at most worker[j]).
// Every worker can be assigned at most one job, but one job can be completed multiple times.
//
// For example, if three workers attempt the same job that pays $1, then the total profit will be $3.
// If a worker cannot complete any job, their profit is $0.
// Return the maximum profit we can achieve after assigning the workers to the jobs.
//
// Example 1:
//
// Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
// Output: 100
// Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get a profit of [20,20,30,30] separately.
//
// Example 2:
//
// Input: difficulty = [85,47,57], profit = [24,66,99], worker = [40,25,25]
// Output: 0
package main

import (
	"fmt"
	"slices"
)

var cases = []struct {
	difficulty []int
	profit     []int
	worker     []int
}{
	{
		difficulty: []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63},
		profit:     []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77},
		worker:     []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16},
	},
	{
		difficulty: []int{13, 37, 58},
		profit:     []int{4, 90, 96},
		worker:     []int{34, 73, 45},
	},
	{
		difficulty: []int{2, 4, 6, 8, 10},
		profit:     []int{10, 20, 30, 40, 50},
		worker:     []int{4, 5, 6, 7},
	},
	{
		difficulty: []int{85, 47, 57},
		profit:     []int{24, 66, 99},
		worker:     []int{40, 25, 25},
	},
}

func main() {
	for i, c := range cases {
		result := maxProfitAssignment(c.difficulty, c.profit, c.worker)

		fmt.Printf("Test case %d; input: maxProfitAssignment(%#v, %#v, %#v); output: %d\n", i, c.difficulty, c.profit, c.worker, result)
	}
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	maxDif := slices.Max(difficulty) + 1
	formatted := make([]int, maxDif)

	difs := map[int]int{}

	for pos, dif := range difficulty {
		pr := profit[pos]

		if val, ok := difs[dif]; ok && pr < val {
			pr = val
		}

		difs[dif] = pr
	}

	if val, ok := difs[0]; ok {
		formatted[0] = val
	} else {
		formatted[0] = 0
	}

	for i := 1; i < maxDif; i++ {
		if val, ok := difs[i]; ok {
			if val > formatted[i-1] {
				formatted[i] = val
			} else {
				formatted[i] = formatted[i-1]
			}
		} else {
			formatted[i] = formatted[i-1]
		}
	}

	sum := 0

	for _, val := range worker {
		if val >= maxDif {
			val = len(formatted) - 1
		}
		sum += formatted[val]
	}

	return sum
}
