package main
func missingNumber(nums []int) int {
	var sum = len(nums)
	//for i:=0;i<len(nums);i++{
	//	sum += i
	//	sum -= nums[i]
	//}
	for i:=0;i<len(nums);i++{
		sum ^= i^nums[i]
	}
	return sum

}
func main(){
	nums := []int{3,0,1,2,4,6}
	print(missingNumber(nums))
}