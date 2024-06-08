package main

import "testing"

func TestExampleFunc(t *testing.T) {
	testCases := []struct {
		input  []int
		output bool
	}{
		{
			input: []int{1, 2, 3},
		},
	}

	for _, test := range testCases {
		result := exampleFunc(test.input)

		if result != test.output {
			t.Errorf("exampleFunc(%#v) = %t; wants %t", test.input, result, test.output)
		}
	}
}
