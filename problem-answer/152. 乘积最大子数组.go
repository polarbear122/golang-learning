package main

import "fmt"

func main() {
	nums := []int{0, 9, 9, 0, 9, 8, 0, -2}
	fmt.Println(maxProduct0(nums))
}

func maxProduct(nums []int) int {
	maxNum := -10
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]*nums[i+1] > nums[i] {
			nums[i+1] = nums[i] * nums[i+1]
		}
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}
	return maxNum
}

func maxProduct0(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max0(mx*nums[i], max0(nums[i], mn*nums[i]))
		minF = min0(mn*nums[i], min0(nums[i], mx*nums[i]))
		ans = max0(maxF, ans)
	}
	return ans
}

func max0(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min0(x, y int) int {
	if x < y {
		return x
	}
	return y
}
