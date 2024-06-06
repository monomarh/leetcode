package main

import "testing"

func TestCountStudents(t *testing.T) {
	tests := []struct {
		students   []int
		sandwiches []int
		result     int
	}{
		{
			students:   []int{1, 1, 0, 0},
			sandwiches: []int{0, 1, 0, 1},
			result:     0,
		},
		{
			students:   []int{1, 1, 1, 0, 0, 1},
			sandwiches: []int{1, 0, 0, 0, 1, 1},
			result:     3,
		},
	}

	for _, test := range tests {
		result := countStudents(test.students, test.sandwiches)
		if result != test.result {
			t.Errorf("countStudents(%v, %v) = %d; want %d", test.students, test.sandwiches, result, test.result)
		}
	}
}
