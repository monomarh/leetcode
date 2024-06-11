package main

import (
	"testing"
)

func TestExampleFunc(t *testing.T) {
	testCases := []struct {
		arr1   []int
		arr2   []int
		sorted []int
	}{
		{
			arr1:   []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			arr2:   []int{2, 1, 4, 3, 9, 6},
			sorted: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			arr1:   []int{28, 6, 22, 8, 44, 17},
			arr2:   []int{22, 28, 8, 6},
			sorted: []int{22, 28, 8, 6, 17, 44},
		},
	}

	for _, test := range testCases {
		result := relativeSortArray(test.arr1, test.arr2)

		if !equalSlices(result, test.sorted) {
			t.Errorf("relativeSortArray(%#v, %#v) = %#v; wants %#v", test.arr1, test.arr2, result, test.sorted)
		}
	}
}

func equalSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}

	return true
}
