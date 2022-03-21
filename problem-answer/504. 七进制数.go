package main

import "strconv"

//给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
func convertToBase7(num int) string {
	if num == 0{
		return "0"
	}
	f := -1
	res := ""
	if num < 0{
		num = num * f
		f = 1
	}
	for num > 0{
		res = strconv.Itoa(num % 7) + res
		num /= 7
	}
	if f == 1{
		res = "-" + res
	}
	return res
}
func main()  {
	num:=-7
	println(convertToBase7(num))
}