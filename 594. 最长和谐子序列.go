package main

import "sort"

//和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。
//现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。
//数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。
func findLHS(nums []int) int {
	sort.Ints(nums)
	res := 0
	start := 0
	for end := 0 ; end < len(nums) ; end++{
		if nums[end] - nums[start] > 1{
			start++
		}
		if nums[end] - nums[start] == 1{
			if res < end-start+1{
				res = end-start+1}
		}
	}
	return res
}

func main()  {
	nums:=[]int{1,3,2,2,5,2,3,7}
	println(findLHS(nums))
}