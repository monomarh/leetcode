package main

import "testing"

func TestRemoveKdigits(t *testing.T) {
	tests := []struct {
		input  string
		k      int
		output string
	}{
		{
			input:  "1432219",
			k:      3,
			output: "1219",
		},
		{
			input:  "10200",
			k:      1,
			output: "200",
		},
		{
			input:  "10",
			k:      2,
			output: "0",
		},
		{
			input:  "112",
			k:      1,
			output: "11",
		},
	}

	for _, test := range tests {
		result := removeKdigits(test.input, test.k)
		if result != test.output {
			t.Errorf("removeKdigits(%s, %d) = %s; want %s", test.input, test.k, result, test.output)
		}
	}
}
