package main

import "testing"

func TestMaxDistance(t *testing.T) {
	testCases := []struct {
		position []int
		m        int
		output   int
	}{
		{
			position: []int{1, 2, 3, 4, 7},
			m:        3,
			output:   3,
		},
		{
			position: []int{5, 4, 3, 2, 1, 1000000000},
			m:        2,
			output:   999999999,
		},
		{
			position: []int{79, 74, 57, 22},
			m:        4,
			output:   5,
		},
	}

	for _, test := range testCases {
		result := maxDistance(test.position, test.m)

		if result != test.output {
			t.Errorf("maxDistance(%#v, %d) = %d; wants %d", test.position, test.m, result, test.output)
		}
	}
}
