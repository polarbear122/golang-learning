package main

//给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
//请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
//任何误差小于 1e-5 的答案都将被视为正确答案。
func findMaxAverage(nums []int, k int) float64 {
	//计算前k个数的和
	initialSum:=0
	for i:=0;i<k;i++ {
		initialSum += nums[i]
	}
	sum:=initialSum
	tempSum:=initialSum
	for i:=0;i<len(nums)-k;i++{
		tempSum-=nums[i]
		tempSum+=nums[i+k]
		if tempSum>sum{
			sum=tempSum
		}
	}
	return float64(sum)/float64(k)
}
func main(){
	nums := []int{9,7,3,5,6,2,0,8,1,9}
	k:=6
	println(findMaxAverage(nums,k))
}
