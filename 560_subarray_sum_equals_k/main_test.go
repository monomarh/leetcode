package main

import "testing"

func TestExampleFunc(t *testing.T) {
	testCases := []struct {
		nums   []int
		k      int
		output int
	}{
		{
			nums:   []int{1, 1, 1},
			k:      2,
			output: 2,
		},
		{
			nums:   []int{1, 2, 3},
			k:      3,
			output: 2,
		},
	}

	for _, test := range testCases {
		result := subarraySum(test.nums, test.k)

		if result != test.output {
			t.Errorf("exampleFunc(%#v, %d) = %d; wants %d", test.nums, test.k, result, test.output)
		}
	}
}
