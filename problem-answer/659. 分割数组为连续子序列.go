package main
func isPossible(nums []int) bool {
	lenNums:=len(nums)
	k:=0
	for i:=0;i<len(nums)-lenNums/3;i++{
		k=1
		for j:=1;j<lenNums/3+1;j++{
			if nums[i+j]!=nums[i]{
				k=0
				break
			}
		}
		if k==1{
			return false
		}
	}
	return true
}
func main()  {
	nums:=[]int{1,2,2,4,5}
	println(isPossible(nums))
}
