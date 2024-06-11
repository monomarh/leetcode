package main

import "testing"

func TestExampleFunc(t *testing.T) {
	testCases := []struct {
		heights    []int
		mismatches int
	}{
		{
			heights:    []int{1, 1, 4, 2, 1, 3},
			mismatches: 3,
		},
		{
			heights:    []int{5, 1, 2, 3, 4},
			mismatches: 5,
		},
		{
			heights:    []int{1, 2, 3, 4, 5},
			mismatches: 0,
		},
	}

	for _, test := range testCases {
		mismatches := heightChecker(test.heights)

		if mismatches != test.mismatches {
			t.Errorf("heightChecker(%#v) = %d; wants %d", test.heights, mismatches, test.mismatches)
		}
	}
}
