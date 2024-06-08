package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2 ListNode

	l1 = *createList([]int{2, 4, 9})
	l2 = *createList([]int{5, 6, 4, 9})

	fmt.Println(addTwoNumbers(&l1, &l2))
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

func createList(arr []int) *ListNode {
	var list *ListNode

	for pos, val := range arr {
		if pos == 0 {
			list = &ListNode{val, nil}
		} else {
			list = &ListNode{val, list}
		}
	}

	return list
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1, num2 = toNumber(l1), toNumber(l2)

	fmt.Println(num1, num2)

	sum := toArray(num1 + num2)

	result := createList(sum)

	for i := result; i != nil; i = i.Next {
		fmt.Println(i)
	}

	return result
}

func toArray(num int) (arr []int) {
	if num == 0 {
		return []int{0}
	}

	for num > 0 {
		arr = append([]int{num % 10}, arr...)
		num = num / 10
	}
	return
}
