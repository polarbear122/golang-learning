package main

import "fmt"

func main() {
	testStr := "  the sky is   blue  "
	fmt.Println(newReverseWords(testStr))
}

func newReverseWords(testStr string) string {
	byteStr := []byte(testStr)
	//去除多余空格
	slow, fast := 0, 0
	for ; byteStr[fast] == ' ' && fast < len(byteStr); fast++ {
		slow++
	}
	byteStr = byteStr[slow:]
	fmt.Println(string(byteStr))
	for slow, fast = 0, 0; slow <= fast && fast < len(byteStr); fast++ {
		if fast > 0 && byteStr[fast] == ' ' && byteStr[fast-1] == ' ' {
			continue
		}
		byteStr[slow] = byteStr[fast]
		slow++
	}
	byteStr = byteStr[:slow]

	if byteStr[len(byteStr)-1] == ' ' {
		byteStr = byteStr[:len(byteStr)-1]
	}
	//反转字符串
	byteStr = byteReverse(byteStr, 0, len(byteStr)-1)

	i := 0
	for i < len(byteStr) {
		j := i
		for ; j < len(byteStr) && byteStr[j] != ' '; j++ {
		}
		byteReverse(byteStr, i, j-1)
		fmt.Println("reverse:", string(byteStr))
		i = j
		i++
	}
	return string(byteStr)
}

func byteReverse(byteStr []byte, begin int, end int) []byte {
	for i := begin; i <= (end+begin)/2; i++ {
		byteStr[i], byteStr[begin+end-i] = byteStr[begin+end-i], byteStr[i]
	}
	return byteStr
}
