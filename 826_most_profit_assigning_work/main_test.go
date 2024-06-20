package main

import "testing"

func TestExampleFunc(t *testing.T) {
	testCases := []struct {
		difficulty []int
		profit     []int
		worker     []int
		output     int
	}{
		{
			difficulty: []int{2, 4, 6, 8, 10},
			profit:     []int{10, 20, 30, 40, 50},
			worker:     []int{4, 5, 6, 7},
			output:     100,
		},
		{
			difficulty: []int{85, 47, 57},
			profit:     []int{24, 66, 99},
			worker:     []int{40, 25, 25},
			output:     0,
		},
	}

	for _, test := range testCases {
		result := maxProfitAssignment(test.difficulty, test.profit, test.worker)

		if result != test.output {
			t.Errorf("maxProfitAssignment(%#v, %#v, %#v) = %d; wants %d", test.difficulty, test.profit, test.worker, result, test.output)
		}
	}
}
