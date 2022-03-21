package main

import "sort"

//给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	result := nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3]
	if nums[0]*nums[1]*nums[len(nums)-1]>result{
		result=nums[0]*nums[1]*nums[len(nums)-1]
	}
	return result
}
func main()  {
	nums:=[]int{1,2,3,0,-1,-99,-98,100,100,100}
	println(maximumProduct(nums))
}