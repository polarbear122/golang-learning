package main

import "fmt"

func majorityElement(nums []int) int {
	num_map:=make(map[int]int,len(nums))

	for _,v := range nums{
		if _,has := num_map[v];!has{
			num_map[v]=1
		}else {
			num_map[v]++
			if num_map[v]>=len(nums)/2{
				return v
			}
		}
	}
	return 0
}

func main() {
	nums:=[]int{2,3,2,1,1,1,1,1,1,22,3,4}
	fmt.Println(majorityElement(nums))
}
