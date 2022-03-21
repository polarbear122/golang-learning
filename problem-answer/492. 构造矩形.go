package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) (result[]int) {
	x:=int(math.Sqrt(float64(area)))
	for i:=x;;i--{
		y:=area/i
		if area%y==0{
			return []int{y, i}
		}
	}
}
func main()  {
	//println(constructRectangle(100))
	fmt.Printf("%v",constructRectangle(105))
}
