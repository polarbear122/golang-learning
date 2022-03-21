package main

import "fmt"

func main() {
	str := "**|**|***|"
	queries := [][]int{{2, 5}, {5, 9}}
	fmt.Println(platesBetweenCandles(str, queries))
}
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	preSum := make([]int, n)
	left := make([]int, n)
	sum, l := 0, -1
	for i, ch := range s {
		if ch == '*' {
			sum++
		} else {
			l = i
		}
		preSum[i] = sum
		left[i] = l
	}

	right := make([]int, n)
	for i, r := n-1, -1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		x, y := right[q[0]], left[q[1]]
		if x >= 0 && y >= 0 && x < y {
			ans[i] = preSum[y] - preSum[x]
		}
	}
	return ans
}

//func platesBetweenCandles(s string, queries [][]int) (result []int) {
//	for _, j := range queries {
//		result = append(result, betweenCandles(s[j[0]:j[1]+1]))
//	}
//	return result
//}
//func betweenCandles(s string) int {
//	var count, tempCount int
//	var candleCount int
//	for i := 0; i < len(s); i++ {
//		if s[i] == '|' {
//			candleCount++
//			count += tempCount
//			tempCount = 0
//
//		}
//		if candleCount > 0 && s[i] == '*' {
//			tempCount++
//		}
//	}
//	return count
//}
