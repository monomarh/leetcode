// Problem description
package main

import "fmt"

var cases = []struct {
	flowerbed []int
	n         int
}{
	{
		flowerbed: []int{0, 0, 0, 0, 0, 1, 0, 0},
		n:         0,
	},
	{
		flowerbed: []int{1, 0, 0, 0, 1},
		n:         1,
	},
	{
		flowerbed: []int{1, 0, 0, 0, 1},
		n:         2,
	},
	{
		flowerbed: []int{0, 0, 1, 0, 0},
		n:         1,
	},
}

func main() {
	for i, c := range cases {
		result := canPlaceFlowers(c.flowerbed, c.n)

		fmt.Printf("Test case %d; input: (%#v, %d); output: %t\n", i, c.flowerbed, c.n, result)
	}
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)

	for pos, val := range flowerbed {
		if val == 0 {
			left := pos == 0 || flowerbed[pos-1] == 0
			right := pos == length-1 || flowerbed[pos+1] == 0

			if left && right {
				n--
				flowerbed[pos] = 1
			}
		}

		if n <= 0 {
			return true
		}
	}

	return n <= 0
}
