package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}

func minRemoveToMakeValid(s string) string {
	checked := make(map[int]bool)

	for pos, val := range s {
		if val == '(' {
			for clPos := pos + 1; clPos < len(s); clPos++ {
				if _, ok := checked[clPos]; s[clPos] == ')' && !ok {
					checked[clPos] = true
					checked[pos] = true
					break
				}
			}
		} else if val != ')' {
			checked[pos] = true
		}
	}

	var buffer bytes.Buffer

	for pos, val := range s {
		if _, ok := checked[pos]; ok {
			buffer.WriteRune(val)
		}
	}

	return buffer.String()
}
