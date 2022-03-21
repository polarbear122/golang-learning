package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	doc := make(map[int]int, len(nums))
	for i, l := 0, len(nums); i < l; i++ {
		if j, ok := doc[nums[i]]; ok && i-j <= k {
			fmt.Println(doc[nums[i]])
			return true
		}
		doc[nums[i]] = i
	}
	return false
}


func main() {
	nums:=[]int{1,3,45,6,6}
	k:=2
	containsNearbyDuplicate(nums,k)
}
