package main

import "testing"

func TestMinRemoveToMakeValid(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "a)b(c)d",
			expected: "ab(c)d",
		},
		{
			input:    "))((",
			expected: "",
		},
	}

	for _, test := range tests {
		result := minRemoveToMakeValid(test.input)
		if result != test.expected {
			t.Errorf("minRemoveToMakeValid(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
