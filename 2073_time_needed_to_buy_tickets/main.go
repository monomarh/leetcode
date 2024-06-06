package main

import "fmt"

func main() {
	fmt.Println(timeRequiredToBuy([]int{2, 3, 2}, 2))
}

// func timeRequiredToBuy(tickets []int, k int) int {
// 	var seconds int

// 	for i := tickets[k]; i >= 0; i-- {
// 		for pos, val := range tickets {
// 			if val != 0 {
// 				tickets[pos]--
// 				seconds++
// 			}
// 			if pos == i && tickets[k] == 0 {
// 				return seconds
// 			}
// 		}
// 	}

// 	return seconds
// }

func timeRequiredToBuy(tickets []int, k int) (seconds int) {
	bilets := tickets[k]

	for pos, val := range tickets {
		if pos <= k {
			seconds += min(bilets, val)
		} else {
			seconds += min(bilets-1, val)
		}
	}

	return
}
