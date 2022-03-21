package main
func maxArea(height []int) int {
	left,right:=0,len(height)-1
	result:=0
	for left<right{
		tempResult:=min(height[left],height[right])*(right-left)
		if tempResult>result{
			result=tempResult
		}
		if height[left]<height[right]{
			left++
		}else {
			right--
		}
	}
	return result
}
func main(){
	height:=[]int{4,3,2,1,4}
	println(maxArea(height))
}
