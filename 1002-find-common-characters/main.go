package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}))
}

func commonChars(words []string) []string {
	arr := make(map[rune]int)
	answer := []string{}

	for _, word := range words {
		for _, value := range word {
			arr[value]++
		}
	}

	for value, count := range arr {
		if count == len(words) {
			answer = append(answer, string(value))
			continue
		}

		if count/len(words) != 0 {
			for i := count / len(words); i > 0; i-- {
				answer = append(answer, string(value))
			}
		}
	}

	return answer
}
