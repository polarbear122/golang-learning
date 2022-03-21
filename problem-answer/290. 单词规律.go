package main


//给定一种规律 pattern和一个字符串str，判断 str 是否遵循相同的规律。
//这里的遵循指完全匹配，例如，pattern里的每个字母和字符串str中的每个非空单词之间存在着双向连接的对应规律。

func wordPattern(pattern string, s string) bool {

	return true

}
func main()  {
	pattern := "abba"
	str := "dog cat cat dog"
	println(wordPattern(pattern,str))
}