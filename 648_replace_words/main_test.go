package main

import "testing"

func TestReplaceWords(t *testing.T) {
	tests := []struct {
		dictionary []string
		sentence   string
		output     string
	}{
		{
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "the cattle was rattled by the battery",
			output:     "the cat was rat by the bat",
		},
		{
			dictionary: []string{"a", "b", "c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
			output:     "a a b c",
		},
	}

	for _, test := range tests {
		result := replaceWords(test.dictionary, test.sentence)

		if result != test.output {
			t.Errorf("replaceWords(%v, %s) = %v, wants %v", test.dictionary, test.sentence, result, test.output)
		}
	}
}
