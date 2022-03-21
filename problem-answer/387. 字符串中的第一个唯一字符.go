package main
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
func firstUniqChar(s string) int {
	k := 0
	for i:=0;i<len(s);i++{
		k = 1
		for j:=0;j<i;j++{
			if s[j] == s[i]{
				k = 0
			}
		}
		for j:=i+1;j<len(s);j++{
			if s[j] == s[i]{
				k = 0
			}
		}
		if k ==1{
			return i
		}
	}
	return -1
}
func main(){
	s := "aabb"
	println(firstUniqChar(s))
}
