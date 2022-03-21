package main

import "strings"

//给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
func repeatedSubstringPattern(s string) bool {
	str2:=s+s
	str2Reduce :=str2[1:len(str2)-1]
	if strings.Contains(str2Reduce,s){
		return true
	}else {
		return false
	}
}
func main()  {
	s:="abab"
	println(repeatedSubstringPattern(s))
}