package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
}

func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)

	for len(hand) > 0 {
		arr := append([]int{}, hand[:1]...)
		hand = hand[1:]

		for j := 0; j < groupSize-1; j++ {
			pos, err := findNextWithPos(hand, arr[len(arr)-1])
			if err != nil {
				return false
			}

			arr = append(arr, hand[pos:pos+1]...)
			hand = append(hand[:pos], hand[pos+1:]...)
		}

		if len(arr) != groupSize {
			return false
		}
	}

	return len(hand) == 0
}

func findNextWithPos(arr []int, previous int) (int, error) {
	for pos, val := range arr {
		if val == previous+1 {
			return pos, nil
		}
	}

	return 0, errors.New("value not found")
}
