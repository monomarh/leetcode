package main

import "testing"

func TestCountSubarrays(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		minK   int
		maxK   int
		result int64
	}{
		{
			name:   "Test case 1",
			nums:   []int{35054, 398719, 945315, 945315, 820417, 945315, 35053, 945315, 171832, 945315, 35054, 109750, 790964, 441974, 552913},
			minK:   35054,
			maxK:   945315,
			result: 19,
		},
		{
			name:   "Test case 2",
			nums:   []int{1, 3, 5, 2, 7, 5},
			minK:   1,
			maxK:   5,
			result: 2,
		},
		{
			name:   "Test case 3",
			nums:   []int{1, 2, 3, 4, 5},
			minK:   6,
			maxK:   4,
			result: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := countSubarrays(test.nums, test.minK, test.maxK)
			if result != test.result {
				t.Errorf("countSubarrays(%v, %d, %d) = %d; want %d", test.nums, test.minK, test.maxK, result, test.result)
			}
		})
	}
}
