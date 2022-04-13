package main

import "fmt"

func main() {
	result := 0
	inputStr := ""
	instr, _ := fmt.Scan(&inputStr)
	fmt.Println("instr:", instr)
	scoreMap := map[int32]int{
		'}': 3,
		']': 2,
		')': 1,
	}
	charMap := map[int32]int32{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	var stack = []int32{0}
	for i, j := range inputStr {
		//fmt.Println(string(stack))
		if i != 0 && j == charMap[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			result += scoreMap[j]
			//fmt.Println("stack:", string(stack), string(j))
		} else {
			stack = append(stack, j)
		}
	}
	fmt.Println(result)
}
