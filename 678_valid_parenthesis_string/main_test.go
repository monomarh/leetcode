package main

import "testing"

func TestCheckValidString(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{
			input: "()",
			valid: true,
		},
		{
			input: "(*)",
			valid: true,
		},
		{
			input: "(*))",
			valid: true,
		},
		{
			input: "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()",
			valid: true,
		},
	}

	for _, test := range tests {
		result := checkValidString(test.input)
		if result != test.valid {
			t.Errorf("checkValidString(\"%s\") = %t; want %t", test.input, result, test.valid)
		}
	}
}
