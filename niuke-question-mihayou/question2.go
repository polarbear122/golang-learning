//米小游的01小串
package main

import "fmt"

func main() {
	inputStr := "1001"
	_, _ = fmt.Scan(&inputStr)
	if inputStr == "" {
		return
	}
	fmt.Println(inputStr)
	result := 0
	for slow, fast := 0, 0; slow < len(inputStr) && fast < len(inputStr); fast++ {
		if inputStr[slow] == '0' && inputStr[fast] == '0' {
			slow++
		} else if inputStr[slow] == '1' && inputStr[fast] == '1' && (fast == len(inputStr)-1 || inputStr[fast+1] == '0') && (fast-slow)%2 == 0 {
			result++
			slow = fast + 1
		}
	}
	fmt.Println(result)
}
