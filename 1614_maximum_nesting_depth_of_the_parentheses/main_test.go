package main

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "(1+(2*3)+((8)/4))+1",
			output: 3,
		},
		{
			input:  "(1)+((2))+(((3)))",
			output: 3,
		},
	}

	for _, test := range tests {
		output := maxDepth(test.input)
		if output != test.output {
			t.Errorf("maxDepth(%s) = %d; wants %d", test.input, output, test.output)
		}
	}
}
