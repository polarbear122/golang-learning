package main

import "math/bits"

//两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。
//给你两个整数 x 和 y，计算并返回它们之间的汉明距离。
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
func main(){
	x,y:=1234,56880
	println(hammingDistance(x,y))
}