package main

import "fmt"

func main() {
	security := []int{5, 3, 3, 3, 5, 6, 2}
	fmt.Println(goodDaysToRobBank(security, 2))
}

func goodDaysToRobBank(security []int, time int) (result []int) {
	n := len(security)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n-1; i++ {
		if security[i] >= security[i+1] {
			left[i+1] = left[i] + 1
		}
		if security[n-i-2] <= security[n-i-1] {
			right[n-i-2] = right[n-i-1] + 1
		}
	}
	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			result = append(result, i)
		}
	}
	return result
}

func increase(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func decrease(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}
