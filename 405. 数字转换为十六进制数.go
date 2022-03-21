package main

import "strconv"

func toHex(num int) (result string) {
	if num == 0 {
		return "0"
	}
	helper := "0123456789abcdef"
	ans := ""
	for num != 0 && len(ans) < 8 {
		temp := num & 15
		ans = string(helper[temp]) + ans
		num >>= 4
	}
	return ans
}
func main(){
	num := 1010
	println(toHex(num))
	var v int64 = 1010
	s16 := strconv.FormatInt(v, 16)
	println(s16)
}
