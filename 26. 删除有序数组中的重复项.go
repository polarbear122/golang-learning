package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums)<=1{
		return len(nums)
	}
	left,right:=1,1
	for left<len(nums)&&right<len(nums){
		if nums[right]!=nums[right-1]{
			nums[left]=nums[right]
			right++
			left++
		}else {
			right++
		}
	}
	return left
}
func main() {
	var nums=[]int{1,1,2,3,3,3,3,4,4}
	fmt.Println(removeDuplicates(nums))
}
