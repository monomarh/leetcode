package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "III",
			output: 3,
		},
		{
			input:  "LVIII",
			output: 58,
		},
		{
			input:  "MCMXCIV",
			output: 1994,
		},
		{
			input:  "XV",
			output: 15,
		},
		{
			input:  "LX",
			output: 60,
		},
	}

	for _, test := range tests {
		result := romanToInt(test.input)
		if result != test.output {
			t.Errorf("romanToInt(\"%s\") = %d; want %d", test.input, result, test.output)
		}
	}
}
