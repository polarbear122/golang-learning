package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	if nums==nil||len(nums)<3{
		return ans
	}
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		if nums[i]>0{
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left,right:=i+1,len(nums)-1
		for left<right{
			sum := nums[i] + nums[left] + nums[right]
			if sum==0{
				res := []int{nums[i], nums[left], nums[right]}
				ans = append(ans, res)
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right  && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}else if sum<0{
				left++
			}else  {
				right--
			}
		}
	}
	return ans
}
func main(){
	nums:=[]int{-10,-2,-4,-5,-4,76,-1,0,1,2,3,5,8,9}
	fmt.Printf("%v",threeSum(nums))
}
