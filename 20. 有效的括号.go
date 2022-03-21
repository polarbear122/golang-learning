package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	stack := make([]byte, 0)
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	for i := 0; i < n; i++ {
		fmt.Println(pairs[s[i]])
		if pairs[s[i]] == 0 {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

func main() {
	str := "{{}{}{}(())}"
	fmt.Println(isValid(str))
}
