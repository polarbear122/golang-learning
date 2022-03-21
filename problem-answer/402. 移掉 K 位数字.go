package main
//给你一个以字符串表示的非负整数 num 和一个整数 k ，
//移除这个数中的 k 位数字，使得剩下的数字最小。
//请你以字符串形式返回这个最小的数字。
//1 <= k <= num.length <= 105
//num 仅由若干位数字（0 - 9）组成
//除了 0 本身之外，num 不含任何前导零
func removeKdigits(num string, k int) string {
	if k == len(num) || num == "0"{
		return "0"
	}
	return "0"
}
func main(){
	num := "10200"
	k := 1
	println(removeKdigits(num,k))
}
