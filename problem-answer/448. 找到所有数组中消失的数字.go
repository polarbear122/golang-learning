package main

import "fmt"

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。
//请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
func findDisappearedNumbers(nums []int) []int {
	t:=0
	result := make([]int,len(nums))
	for i:=1;i<len(nums)+1;i++{
		k:=1
		for j:=0;j<len(nums);j++{
			if i==nums[j]{
				k=0
			}
		}
		if k==1{
			result[t]=i
			t++
		}
	}
	return result[0:t]
}
func main(){
	n :=[]int{4,3,2,7,8,2,3,1}
	fmt.Printf("%v", findDisappearedNumbers(n))
}
