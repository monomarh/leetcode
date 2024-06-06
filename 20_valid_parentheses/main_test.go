package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "(){}[]",
			output: true,
		},
		{
			input:  "({})",
			output: true,
		},
		{
			input:  "(})",
			output: false,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			input:  "({{{{}}}))",
			output: false,
		},
		{
			input:  "{(})",
			output: false,
		},
		{
			input:  "(){[{()}]}",
			output: true,
		},
		{
			input:  "{[(]})",
			output: false,
		},
	}

	for _, test := range tests {
		result := isValid(test.input)
		if result != test.output {
			t.Errorf("isValid(\"%s\") = %t; want %t", test.input, result, test.output)
		}
	}
}
