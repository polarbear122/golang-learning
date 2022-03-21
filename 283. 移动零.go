package main
//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
func moveZeroes(nums []int)  {
	var left,right,n = 0,0,len(nums)
	for right <n{
		if nums[right] != 0{
			nums[left],nums[right] = nums[right],nums[left]
			left ++
		}
		right++
	}
}
func main(){
	nums := []int{0,0,1}
	moveZeroes(nums)
	for i:=0;i<len(nums);i++{
		println(nums[i])
	}
}