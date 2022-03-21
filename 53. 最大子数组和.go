package main

import "fmt"

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//子数组 是数组中的一个连续部分。
func maxSubArray(nums []int) int {
	max:=nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i]+nums[i-1]>nums[i]{
			nums[i]+=nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
func main() {
	nums:=[]int{1,2,-3,4,5,6,6,7,8,9,10}
	fmt.Println(maxSubArray(nums))
}
