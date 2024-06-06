package main

import (
	"fmt"
	"sort"
)

func main() {
	// cards := []int{2, 17}
	cards := []int{17, 13, 11, 2, 3, 5, 7}

	fmt.Println(deckRevealedIncreasing(cards))
}

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	var queue []int

	for i := 0; i < len(deck); i++ {
		queue = append(queue, i)
	}

	result := make([]int, len(deck))

	for _, card := range deck {
		result[queue[0]] = card
		queue = queue[1:]
		if len(queue) > 0 {
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
	}

	return result
}
