package main

import "math"

func divide(a int, b int) int {
	if a == math.MaxInt32 && b == -1 {
		return math.MaxInt32
	}
	sign := 1
	if sign > 0 {
		for i := 0; i < a; i++ {

		}
	}
	return sign
}
func main() {
	println(divide(10, 2))
}
