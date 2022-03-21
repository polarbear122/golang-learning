package main

import "fmt"

func climbStairs(n int) int {
	var x,y,z =0,0,1
	for i:=0;i<n;i++{
		x=y
		y=z
		z=x+y
	}
	return z
}
func main() {
	n:=10
	fmt.Println(climbStairs(n))
}
