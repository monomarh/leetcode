package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func toNumber(l *ListNode) (num int) {
	arr := []int{}
	for i := l; i != nil; i = i.Next {
		arr = append(arr, i.Val)
	}
	zeros := 0
	for i := len(arr) - 1; i >= 0; i-- {
		num += arr[i] * int(math.Pow(10, float64(zeros)))
		zeros++
	}
	return
}

func main() {
	var l1, l2 ListNode

	l1 = *createList([]int{2, 4, 9})
	l2 = *createList([]int{5, 6, 4, 9})

	fmt.Println(addTwoNumbers(&l1, &l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1, num2 = toNumber(l1), toNumber(l2)

	sum := toArray(num1 + num2)

	result := createList(sum)

	return result
}
