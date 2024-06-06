package main

import "testing"

func TestTimeRequiredToBuy(t *testing.T) {
	tests := []struct {
		input    []int
		position int
		output   int
	}{
		{
			input:    []int{2, 3, 2},
			position: 2,
			output:   6,
		},
		{
			input:    []int{5, 1, 1, 1},
			position: 0,
			output:   8,
		},
	}

	for _, test := range tests {
		result := timeRequiredToBuy(test.input, test.position)
		if result != test.output {
			t.Errorf("timeRequiredToBuy(%v, %d) = %d; want %d", test.input, test.position, result, test.output)
		}
	}
}
