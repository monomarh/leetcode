package main

import "fmt"

func main() {
	fmt.Println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}

func countStudents(students []int, sandwiches []int) int {
	circularCount, squareCount := 0, 0
	for _, student := range students {
		if student == 0 {
			circularCount++
		} else {
			squareCount++
		}
	}

	for len(sandwiches) > 0 {
		if sandwiches[0] == 0 {
			if circularCount == 0 {
				break
			}
			sandwiches = sandwiches[1:]
			circularCount--
		} else {
			if squareCount == 0 {
				break
			}
			sandwiches = sandwiches[1:]
			squareCount--
		}
	}

	return circularCount + squareCount
}
