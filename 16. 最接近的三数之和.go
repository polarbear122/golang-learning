package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	totalDiff,tempDiff:=1000000,0
	low1,low2,low3:=0,0,0
	for i:=0;i<len(nums);i++{
		println(i)
		for left,right:=i+1,len(nums)-1;left<right;{
			tempDiff=nums[i]+nums[left]+nums[right]-target
			println("tempDiff",tempDiff,left,right)
			if math.Abs(float64(tempDiff))<math.Abs(float64(totalDiff)){
				totalDiff=tempDiff
				low1,low2,low3=i,left,right
			}
			if tempDiff>0{
				right--
			}else if tempDiff<0{
				left++
			}else{
				return target
			}
		}
	}
	return nums[low1]+nums[low2]+nums[low3]
}
func main(){
	nums := []int{23,-57,-16,2,-22,58,-68,-90,98,18,-66,11,98,41,18,74,-90,58,100,62,32,44,48,34,-16,-5,-74,60,28,54,-65,-53,24,92,4,-10,-39,-68,88,-65,86,-18,-87,-64,-44,94,-77,95,91,-91,-39,63,-55,31,-99,78,-66,9,39,67,56,58,-73,67,-39,23,91,-89,83,-86,-81,64,25,99,-12,14,91,-34,91,-93,-90,-40,-34,-24,-80,63,40,39,-33,-62,-1,46,-14,-22,12,-67,-29,-15,67,46,-49,-72,-69,-20,48,12,-24,75,4,8,-43,-86,-13,-2,-5,-92,-23,-9,75,59,-44,-4,-30,-72,-7,84,88,-56,-1,-47,-72,28,68,-19,33,98,50,-10,-13,-12,-60,70,31,-14,-15,6,65,-15}
	target := -56
	println(threeSumClosest(nums,target))
}
