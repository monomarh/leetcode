package main

import "testing"

func TestExists(t *testing.T) {
	tests := []struct {
		board  [][]byte
		word   string
		exists bool
	}{
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "ABCCED",
			exists: true,
		},
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "SEE",
			exists: true,
		},
		{
			board:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:   "ABCB",
			exists: false,
		},
	}

	for _, test := range tests {
		exists := exist(test.board, test.word)

		if test.exists != exists {
			t.Errorf("exists(%v, %s) = %t; want %t", test.board, test.word, exists, test.exists)
		}
	}
}
