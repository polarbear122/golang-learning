package main

import (
	"fmt"
	"sort"
)

//集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，
//导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。
//给定一个数组 nums 代表了集合 S 发生错误后的结果。
//请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	var repeatAddress,errorAddress =0,-1
	for i:=0;i<len(nums)-1;i++{
		if nums[i]==nums[i+1]{
			repeatAddress=i+1
		}
		if nums[i+1]-nums[i]==2{
			errorAddress=nums[i]+1
		}
	}
	if nums[0]!=1{
		errorAddress=1
	}
	if nums[len(nums)-1]!=len(nums) {
		errorAddress=len(nums)
	}
	var result = []int{nums[repeatAddress],errorAddress}
	return result
}
func main()  {
	nums:=[]int{2,2}
	fmt.Printf("%v",findErrorNums(nums))
}