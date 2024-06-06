package main

import "fmt"

func main() {
	fmt.Println(checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"))
}

func checkValidString(s string) bool {
	open, close, stars := 0, 0, 0
	for _, c := range s {
		switch c {
		case '(':
			open++
		case ')':
			close++
		case '*':
			stars++
		}
		if close > open+stars {
			return false
		}
	}

	open, close, stars = 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			open++
		case ')':
			close++
		case '*':
			stars++
		}
		if open > close+stars {
			return false
		}
	}

	return true
}
