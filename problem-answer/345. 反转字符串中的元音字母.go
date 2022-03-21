package main

import "strings"

//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
//元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。

func reverseVowels(s string) string {
	new_s := []byte(s)
	var left,right = 0,len(new_s )-1
	for left<right{
		for left < len(new_s ) && !strings.Contains("aeiouAEIOU", string(new_s [left])) {
			left++
		}
		for right > 0 && !strings.Contains("aeiouAEIOU", string(new_s [right])){
			right--
		}
		if left<right{
			new_s [left],new_s [right] =new_s [right],new_s [left]
			left++
			right--
		}

	}
	return string(new_s )
}

func main()  {
	s := "hello"
	println(reverseVowels(s))
}