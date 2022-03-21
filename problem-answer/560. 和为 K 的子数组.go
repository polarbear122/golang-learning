package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 1}
	k := 2
	fmt.Println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) (result int) {
	m := make(map[int]int, len(nums))
	pre := 0
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if m[pre-k] != 0 {
			result += m[pre-k]
		}
		m[pre]++
	}
	return
}
