package main
func findMaxConsecutiveOnes(nums []int) int {
	left,result:=0,0
	for i:=0;i<len(nums);{
		if nums[i]!=0{
			left=i
			tempResult:=0
			for j:=left;j<len(nums);j++{
				println("left",left)
				if nums[j]==0{
					i=j
					break
				}
				tempResult++
			}
			if tempResult>result{
				result=tempResult
			}
		}
		i++
	}
	return result
}
func main()  {
	num := []int{1,0,1,1,0,1}
	println(findMaxConsecutiveOnes(num))
}
