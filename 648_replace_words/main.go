// In English, we have a concept called root,
// which can be followed by some other word to form another longer word
// let's call this word derivative.
//
// For example, when the root "help" is followed by the word "ful", we can form a derivative "helpful".
//
// Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces,
// replace all the derivatives in the sentence with the root forming it.
// If a derivative can be replaced by more than one root, replace it with the root that has the shortest length.
//
// Return the sentence after the replacement.
//
// Example 1:
//
// Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// Output: "the cat was rat by the bat"
//
// Example 2:
//
// Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
// Output: "a a b c"
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(replaceWords(
		[]string{"cat", "bat", "rat"},
		"the cattle was rattled by the battery",
	))
}

// Version without using trie
/* func replaceWords(dictionary []string, sentence string) string {
	words := strings.Fields(sentence)

	for pos, word := range words {
		sub := ""

		for _, root := range dictionary {
			isRoot := strings.HasPrefix(word, root)

			if isRoot {
				if sub != "" && len(root) > len(sub) {
					continue
				}

				if sub != "" && strings.Index(word, root) > strings.Index(word, sub) {
					continue
				}

				sub = root
			}
		}

		if sub != "" {
			words[pos] = sub
		}
	}

	return strings.Join(words, " ")
} */

type Node struct {
	end   bool
	chars ['z' - 'a' + 1]*Node
}

func insert(root *Node, word string) {
	for _, val := range word {
		if root.chars[val-'a'] == nil {
			root.chars[val-'a'] = &Node{}
		}
		root = root.chars[val-'a']
	}
	root.end = true
}

func find(root *Node, word string) (string, error) {
	out := make([]rune, 0, len(word))

	for _, val := range word {
		if root.chars[val-'a'] == nil {
			return "", errors.New("substring wasn't found")
		}

		out = append(out, val)
		root = root.chars[val-'a']

		if root.end {
			return string(out), nil
		}
	}

	if root.end {
		return string(out), nil
	}

	return "", errors.New("substring wasn't found")
}

func replaceWords(dictionary []string, sentence string) string {
	root := &Node{}
	for _, word := range dictionary {
		insert(root, word)
	}

	words := strings.Split(sentence, " ")

	for pos, val := range words {
		sub, err := find(root, val)
		if err != nil {
			continue
		}
		words[pos] = sub
	}

	return strings.Join(words, " ")
}
