package main

import "fmt"

func containsDuplicate(nums []int) bool {
	contain:=map[int]bool{}

	for _,num := range nums {
		if !contain[num]{
			contain[num]=true
		}else {
			return true
		}
	}
	return false
}
func main() {
	nums:=[]int{2,54,6,7,8,8}
	fmt.Println(containsDuplicate(nums))
}
