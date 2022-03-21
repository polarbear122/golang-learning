package main

import (
	"fmt"
	"math/rand"
	"time"
)

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {
	st := time.Now()
	fmt.Println(myPow(0.00001, 2147483647))
	fmt.Println("time:", time.Since(st))
	//for i := 0; i < 10; i++ {
	//	x := randFloat()
	//	n := randInt100()
	//
	//	if math.Round(math.Pow(x, float64(n))) == math.Round(myPow(x, n)) {
	//		fmt.Println("success", x, n, myPow(x, n))
	//	} else {
	//		fmt.Println("fail", x, n, myPow(x, n))
	//	}
	//}
}

func randInt100() int {
	return rand.Int() % 100
}
func randFloat() float64 {
	return rand.Float64()
}
