package main

import "fmt"

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Low,nums2Low:=0,0
	lenNum1,lenNum2:=len(nums1),len(nums2)
	count:=0
	newNum:=[]int{}
	//newNum := make([]int,len(nums1)+len(nums2))
	for count<(lenNum1+lenNum2)/2+2{
		if nums1Low>=lenNum1&&nums2Low<lenNum2{
			newNum=append(newNum,nums2[nums2Low])
			nums2Low++
			println("line1")
		}else if nums2Low>=lenNum2&&nums1Low<lenNum1{
			newNum=append(newNum,nums1[nums1Low])
			nums1Low++
			println("line2")
		}else if nums1Low<lenNum1&&nums1[nums1Low]<nums2[nums2Low]{
			newNum=append(newNum,nums1[nums1Low])
			println("line3")
			nums1Low++
		}else if nums2Low<lenNum2&&nums1[nums1Low]>=nums2[nums2Low]{
			newNum=append(newNum,nums2[nums2Low])
			println("line4")
			nums2Low++
		}else{
			println("error",nums1Low,nums2Low)
		}
		count++
	}
	fmt.Printf("%v",newNum)
	if (lenNum1+lenNum2)%2==1{
		return float64(newNum[(lenNum1+lenNum2)/2])
	}else {
		return float64(newNum[(lenNum1+lenNum2)/2]+newNum[(lenNum1+lenNum2)/2-1])/2
	}
}
func main(){
	nums1 := []int{}
	nums2 := []int{2,2,3,4,5,65,67,99}
	println(findMedianSortedArrays(nums1,nums2))
}

