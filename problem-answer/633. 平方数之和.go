package main

import "math"

func judgeSquareSum(c int) bool {
	for i:=0;i*i<c;i++{
		if int(math.Sqrt(float64(c-i*i)))*int(math.Sqrt(float64(c-i*i)))+i*i==c{
			return true
		}
	}
	return false
}
func main()  {
	println(judgeSquareSum(7))
}
