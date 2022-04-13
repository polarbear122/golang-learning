package main

import "fmt"

func main() {
	n, m, k := 0, 0, 0
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&k)

	str := []string{}
	for i := 0; i < n; i++ {
		tmp := ""
		fmt.Scan(&tmp)
		str = append(str, tmp)
	}
	stone := [][]int{}
	for i := 0; i < k; i++ {
		tmpsn := []int{}
		for j := 0; j < 2; j++ {
			tmps := 0
			fmt.Scan(&tmps)
			tmpsn = append(tmpsn, tmps)
		}
		stone = append(stone, tmpsn)
	}
	stoneMap := [][]int{}
	start, end := false, false
	//fmt.Printf("%v,%v\n", str, stone)
	for i := 0; i < len(str); i++ {
		for j, v := range str[i] {
			if stone[i+1][j+1] == 0 {
				continue
			}
			if v == 'S' {
				start = true
			}
			if v == 'F' {
				end = true
			}
			stoneMap[i][j+1] = 1
			stoneMap[i][j-1] = 1
			stoneMap[i+1][j] = 1
			stoneMap[i-1][j] = 1
		}
	}
	if start && end {
		fmt.Println(0)
	}
}
