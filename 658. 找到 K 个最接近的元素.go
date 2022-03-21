package main

import "fmt"

//给定一个排序好的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
//整数 a 比整数 b 更接近 x 需要满足：
//|a - x| < |b - x| 或者
//|a - x| == |b - x| 且 a < b
func findClosestElements(arr []int, k int, x int) []int {
	if len(arr)==k{
		return arr
	}
	//找出x的位置
	xAddress := 0
	var result []int
	for i:=0;i<len(arr);i++{
		if arr[i]<x{
			xAddress = i
		}
	}
	if xAddress<len(arr)-1&&arr[xAddress+1]-x<x-arr[xAddress]{
		xAddress++
	}
	//
	left,right:=xAddress,xAddress
	count:=0
	for{
		count++
		if count==k{
			break
		}
		println(left,right,count)
		if right>=len(arr)-1&&left>0{
			left--
			continue
		}
		if left<=0&&right<len(arr)-1{
			right++
			continue
		}
		if (x-arr[left-1])<=(arr[right+1]-x){
			left--
		}else {
			right++
		}
	}
	println("left,right",left,right)
	result=arr[left:right+1]

	return result
}
func main()  {
	arr:=[]int{-10,-5,0,0,1,1,2,2,3,3,3,3,3,3,5,6,7,111}
	k:=1
	x:=1000
	fmt.Printf("%v",findClosestElements(arr,k,x))
}