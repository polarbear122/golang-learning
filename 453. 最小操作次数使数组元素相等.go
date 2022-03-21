package main

//给定一个长度为 n 的 非空 整数数组，每次操作将会使 n - 1 个元素增加 1。找出让数组所有元素相等的最小操作次数。
func minMoves(nums []int) int {
	if len(nums)==0{
		return 0
	}
	minNums:=nums[0]
	sums:=0
	for i:=0;i<len(nums);i++{
		if nums[i]<minNums{
			minNums=nums[i]
		}
		sums+=nums[i]
	}
	return sums-minNums*len(nums)
}
func main()  {
	nums:=[]int{1,2,3}
	println(minMoves(nums))
}