package main

import "fmt"

func main() {
	str := ""
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
	}
	x, y := 0, 0 //初试位置
	for _, v := range str {
		if v == 'R' {
			x++
		} else if v == 'L' {
			x--
		} else if v == 'U' {
			y++
		} else if v == 'D' {
			y--
		}
	}
	fmt.Printf("(%d,%d)", x, y)
}
