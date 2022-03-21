package main

import "fmt"

func removeElement(nums []int, val int) int {
	left:=0
	for i:=0;i<len(nums);i++{
		if nums[i]!=val{
			nums[left]=nums[i]
			left++
		}
	}
	return left
}
func main() {
	nums:=[]int{1,1,1,2,2,3,4,5,6}
	val:=1
	fmt.Println(removeElement(nums,val))
}
