package main

func firstBadVersion(n int) int {
	var left,right = 0,n
	for i:=0;i<n;i++{
		mid := (left+right)/2
		if isBadVersion(mid) == true {
			left = mid+1
		}else{
			right = mid
		}
		if left == right{
			return left
		}
	}
	return 0
}
func isBadVersion(version int) bool{
	if version <10{
		return true
	}
	return false
}
func main()  {
	n := 10
	println(firstBadVersion(n))
}