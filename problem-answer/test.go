package main

import (
	"fmt"
	"time"
)

func main() {
	//num := [2]int{10, 11}
	num := new([10]int)
	//for i := 0; i < 10; i++ {
	//	num[i] = i
	//}
	num[1] = 1
	nums := []int{10, 11}
	fmt.Println(num, nums)

	time.Sleep(time.Second)
}
