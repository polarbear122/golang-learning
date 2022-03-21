package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(CalculateMethodCount(i))
	}
}

func CalculateMethodCount(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 || num == 2 {
		return num
	} else if num == 3 {
		return 4
	}
	return CalculateMethodCount(num-1) + CalculateMethodCount(num-2) + CalculateMethodCount(num-3)

}
