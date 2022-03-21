package main

import "fmt"

func canJump(nums []int) bool {
	i,rightBox,n:=0,nums[0],len(nums)-1
	for ;i<=rightBox&&i<=n;i++{
		if i+nums[i]>rightBox{
			rightBox=i+nums[i]
		}
	}
	return rightBox>=n
}
func main() {
	nums:=[]int{0}
	fmt.Println(canJump(nums))
}
