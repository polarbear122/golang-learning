package main

import "fmt"

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
func letterCombinations(digits string) (result[]string) {
	var letter = [][]byte{{'a','b','c'},{'d','e','f'},{'g','h','i'},{'j','k','l'},{'m','n','o'},{'p','q','r','s'},{'t','u','v'},{'w','x','y','z'}}
	for _,digit:=range digits{
		for j:=0;j<len(letter[int(digit)-'2']);j++{
			result=append(result,string(letter[int(digit)-'2'][j]))
		}
	}
	return
}
func main(){
	digits := "23"
	fmt.Printf("%v",letterCombinations(digits))
}
