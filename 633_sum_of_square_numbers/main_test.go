package main

import "testing"

func TestExampleFunc(t *testing.T) {
	testCases := []struct {
		input  int
		output bool
	}{
		{
			input:  5,
			output: true,
		},
		{
			input:  3,
			output: false,
		},
		{
			input:  2,
			output: true,
		},
	}

	for _, test := range testCases {
		result := judgeSquareSum(test.input)

		if result != test.output {
			t.Errorf("judgeSquareSum(%d) = %t; wants %t", test.input, result, test.output)
		}
	}
}
