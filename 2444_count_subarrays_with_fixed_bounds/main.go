package main

import (
	"fmt"
)

func main() {
	nums := []int{35054, 398719, 945315, 945315, 820417, 945315, 35053, 945315, 171832, 945315, 35054, 109750, 790964, 441974, 552913}

	fmt.Print(countSubarrays(nums, 35054, 945315))
}

func countSubarrays(nums []int, minK int, maxK int) (res int64) {
	left, pmin, pmax := -1, -1, -1

	for right, num := range nums {
		if num < minK || num > maxK {
			left = right
		}
		if num == minK {
			pmin = right
		}
		if num == maxK {
			pmax = right
		}
		res += int64(max(0, min(pmin, pmax)-left))
	}

	return
}
