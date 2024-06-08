// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
//
// A subarray is a contiguous non-empty sequence of elements within an array.
//
// Example 1:
//
// Input: nums = [1,1,1], k = 2
// Output: 2
//
// Example 2:
//
// Input: nums = [1,2,3], k = 3
// Output: 2
package main

import "fmt"

var cases = []struct {
	nums []int
	k    int
}{
	{
		nums: []int{1, 1, 1},
		k:    2,
	},
	{
		nums: []int{1, 2, 3},
		k:    3,
	},
}

func main() {
	for i, c := range cases {
		result := subarraySum(c.nums, c.k)

		fmt.Printf("Test case %d; input: (%#v, %d); output: %d\n", i, c.nums, c.k, result)
	}
}

func subarraySum(nums []int, k int) int {
	mp := make(map[int]int, len(nums))

	var sum, count int

	mp[0] = 1

	for _, val := range nums {
		sum += val
		count += mp[sum-k]
		mp[sum]++
	}

	return count
}
