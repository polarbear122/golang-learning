package main

import "fmt"

//给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。
//请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。
//如果不存在，对应位置输出 -1 。
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	for i:=0;i<len(nums1);i++{
		k:=1
		for j:=i;j<len(nums2);j++{
			if nums2[j]>nums1[i]{
				k=0
				nums1[i]=nums2[j]
				break}
		}
		if k==1{
			nums1[i]=-1
		}
	}
	return nums1
}
func main()  {
	nums1 := []int{4,1,2}
	nums2 := []int{1,3,4,2}
	fmt.Printf("%v\n",nextGreaterElement(nums1,nums2))
	println(nextGreaterElement(nums1,nums2))
}