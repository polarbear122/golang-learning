package main

//给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
func reverseWords(s string) string {
	t := 0
	m := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			for j := t; j < (i+t)/2; j++ {
				m[j], m[i-j-1+t] = m[i-j-1+t], m[j]
			}
			t = i + 1
		}
	}
	for i := len(s) - 1; i >= (t+len(s))/2; i-- {
		m[i], m[t+len(s)-i-1] = m[t+len(s)-i-1], m[i]
		println("j,i-j-1", i, t+len(s)-i-1)
	}
	return string(m)
}
func main() {
	println(reverseWords("Let's take LeetCode contest"))
}
