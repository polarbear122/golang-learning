package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) (result[][]int) {
	if len(nums )<4{
		return result
	}
	sort.Ints(nums )
	tempDiff:=0
	for i:=0;i<len(nums)-3;i++{
		if i!=0&&nums[i]==nums[i-1]{
			continue}
		for j:=i+1;j<len(nums)-2;j++{
			if j!=i+1&&nums[j]==nums[j-1]{
				continue}
			for left,right:=j+1,len(nums)-1;left<right;{
				if left!=j+1&&nums[left]==nums[left-1]{
					left++
					continue}
				if right!=len(nums)-1 && nums[right]==nums[right+1]{
					right--
					continue}
				tempDiff=nums[i]+nums[j]+nums[left]+nums[right]-target
				if tempDiff==0 {
					tempResult := []int{nums[i],nums[j],nums[left],nums[right]}
					result = append(result, tempResult)
					right--
					left++
				}else if tempDiff>0{
					right--
				}else {
					left++}
				}
			}
		}
	return
}
func main() {
	nums := []int{-4,0,1,2,-1,-1}
	target := -1
	fmt.Printf("%v",fourSum(nums,target))

}
