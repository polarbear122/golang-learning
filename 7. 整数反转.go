package main

import "math"

func reverse(x int) int {
	sign:=1
	if x<0{
		sign=-1
		x=-x
	}
	result:=0
	for x>0{
		result*=10
		result+=x%10
		x/=10
	}
	result*=sign
	if result-int(math.Pow(2,31))>=0||result+int(math.Pow(2,31))<=0{
		result=0
	}
	return result
}
func main()  {
	println(reverse(123444))
	println(reverse(-102))
	println(reverse(0))
	println(reverse(-1))
	println(int(math.Pow(2,31)))
}
