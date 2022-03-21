package main
//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
func isPalindrome(x int) bool {
	if x<0{
		return false
	}else if x==0 {
		return true
	}
	num :=""
	for x>0{
		num+=string(rune(x % 10))
		x/=10
	}
	for i:=0;i<=len(num)/2;i++{
		if num[i]!=num[len(num)-i-1]{
			return false
		}
	}
	return true
}
func main()  {
	println(isPalindrome(12134))
	println(isPalindrome(1))
	println(isPalindrome(101))
}
