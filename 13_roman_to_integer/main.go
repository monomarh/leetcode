package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) (result int) {
	arr := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for pos, val := range s {
		if num := arr[val]; num < arr[rune(s[min(len(s)-1, pos+1)])] {
			result -= num
		} else {
			result += num
		}
	}

	return
}
